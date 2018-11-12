package put

import (
	"context"
	"fmt"
	"github.com/minio/minio/pkg/auth"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"path/filepath"
	"storj/ditto/cmd/utils"
	dcontext "storj/ditto/pkg/context"
	fsystem "storj/ditto/pkg/filesys"
	l "storj/ditto/pkg/logger"
	"storj/ditto/pkg/uploader"
)

type putExec struct {
	resolver utils.GetwayResolver
	uploader.ObjLayerAsyncUploader
	fsystem.FsCheckDir
	logger l.Logger
}

func NewPutExec(resolver utils.GetwayResolver, logger l.Logger) putExec {
	uploader := uploader.NewFolderUploader(nil, fsystem.NewHashFileReader(), fsystem.NewDirReader(), logger)
	return newPutExec(resolver, uploader, fsystem.DirChecker(), logger)
}

func newPutExec(resolver utils.GetwayResolver, uploader uploader.ObjLayerAsyncUploader, dirChecker fsystem.FsCheckDir, logger l.Logger) putExec {
	return putExec{resolver, uploader, dirChecker, logger }
}

func (e putExec) logF(format string, params ...interface{}) {
	e.logger.Log(fmt.Sprintf(format, params...))
}

//Main function
func (e putExec) runE(cmd *cobra.Command, args []string) error {
	sigc = make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)

	gw, err := e.resolver(e.logger)
	if err != nil {
		return err
	}

	mirr, err := gw.NewGatewayLayer(auth.Credentials{})
	if err != nil {
		return err
	}

	e.SetObjLayer(mirr)

	bctx := context.Background()
	_, err = mirr.GetBucketInfo(bctx, args[0])
	if err != nil {
		return err
	}

	isDir, err := e.CheckIfDir(args[1])
	if err != nil {
		return err
	}

	ctx, cancelf := context.WithCancel(bctx)
	defer func() {
		cancelf()
	}()
	
	lpath := filepath.Clean(args[1])

	ctxp := dcontext.NewPutCtx(
		ctx,
		frecursive,
		fforce,
		fprefix,
		fdelimiter,
		lpath)

	var errc <-chan error
	if isDir {
		errc = e.UploadFolderAsync(ctxp, args[0], lpath)
	} else {
		errc = e.UploadFileAsync(ctxp, args[0], lpath)
	}

	tnum := 1
	for i:= 0; i < tnum; i++ {
		select {
		case err = <-errc:
			//e.logger.LogE(err)
		case sig := <-sigc:
			e.logF("Catched interrupt! %s\n", sig)
			cancelf()
			tnum++
		}
	}

	return err
}