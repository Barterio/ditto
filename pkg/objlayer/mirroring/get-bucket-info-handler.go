// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package mirroring

import (
	"context"
	minio "github.com/minio/minio/cmd"
)

func NewGetBucketInfoHandler(m 	     *MirroringObjectLayer,
							 ctx     context.Context,
							 bucket  string) *getBucketInfoHandler {

	h := &getBucketInfoHandler{}

	h.m = m
	h.ctx =  ctx
	h.primeBucket = h.m.getPrimeBucketName(h.m.Config.GetObjectOptions.DefaultOptions.DefaultSource, bucket)
	h.alterBucket = h.m.getAlterBucketName(h.m.Config.GetObjectOptions.DefaultOptions.DefaultSource, bucket)

	return h
}

type getBucketInfoHandler struct {
	baseHandler
	primeBucket string
	alterBucket string
	primeInfo   minio.BucketInfo
	alterInfo   minio.BucketInfo
}

func (h *getBucketInfoHandler) execPrime() *getBucketInfoHandler {
	h.primeInfo, h.primeErr = h.m.Prime.GetBucketInfo(h.ctx, h.primeBucket)

	return h
}

func (h *getBucketInfoHandler) execAlter() *getBucketInfoHandler {
	h.alterInfo, h.alterErr = h.m.Alter.GetBucketInfo(h.ctx, h.alterBucket)

	return h
}

func (h *getBucketInfoHandler) Process () (objInfo minio.BucketInfo, err error) {

	h.execPrime()

	if h.primeErr == nil {
		return h.primeInfo, nil
	}

	// h.m.Logger.LogE(h.primeErr)

	h.execAlter()

	if h.alterErr != nil {

		h.m.Logger.LogE(h.alterErr)

		return h.alterInfo, h.primeErr
	}

	return h.alterInfo, nil
}


