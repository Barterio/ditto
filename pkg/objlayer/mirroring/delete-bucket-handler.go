// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package mirroring

import (
	"context"
)

func NewDeleteBucketHandler(m *MirroringObjectLayer, ctx context.Context, bucket string) *deleteBucketHandler {

	h := &deleteBucketHandler{}

	h.m = m
	h.ctx =  ctx
	h.primeBucket = h.m.getPrimeBucketName(h.m.Config.DeleteOptions.DefaultOptions.DefaultSource, bucket)
	h.alterBucket = h.m.getAlterBucketName(h.m.Config.DeleteOptions.DefaultOptions.DefaultSource, bucket)

	h.m.Prime, h.m.Alter = selectServer(h.m, h.m.Config.DeleteOptions.DefaultOptions.DefaultSource)

	return h
}

type deleteBucketHandler struct {
	baseHandler
	primeBucket string
	alterBucket string
	object string
}

func (h *deleteBucketHandler) execPrime() *deleteBucketHandler {
	h.primeErr = h.m.Prime.DeleteBucket(h.ctx, h.primeBucket)

	return h
}

func (h *deleteBucketHandler) execAlter() *deleteBucketHandler {
	h.alterErr = h.m.Alter.DeleteBucket(h.ctx, h.alterBucket)

	return h
}

func (h *deleteBucketHandler) Process () error {
	h.execPrime()

	if h.primeErr != nil {
		return  h.primeErr
	}

	h.execAlter()

	if h.alterErr != nil {
		//h.m.Logger.Err = h.alterErr
	}

	return nil
}


