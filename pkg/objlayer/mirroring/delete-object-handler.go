// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package mirroring

import (
	"context"
)

func NewDeleteObjectHandler(m *MirroringObjectLayer, ctx context.Context, bucket, object string) *deleteObjectHandler {

	h := &deleteObjectHandler{}

	h.m = m
	h.ctx =  ctx
	h.primeBucket = h.m.getPrimeBucketName(h.m.Config.DeleteOptions.DefaultOptions.DefaultSource, bucket)
	h.alterBucket = h.m.getAlterBucketName(h.m.Config.DeleteOptions.DefaultOptions.DefaultSource, bucket)
	h.object = object

	h.m.Prime, h.m.Alter = selectServer(h.m, h.m.Config.DeleteOptions.DefaultOptions.DefaultSource)

	return h
}

type deleteObjectHandler struct {
	baseHandler
	primeBucket string
	alterBucket string
	object string
}

func (h *deleteObjectHandler) execPrime() *deleteObjectHandler {
	h.primeErr = h.m.Prime.DeleteObject(h.ctx, h.primeBucket, h.object)

	return h
}

func (h *deleteObjectHandler) execAlter() *deleteObjectHandler {
	h.alterErr = h.m.Alter.DeleteObject(h.ctx, h.alterBucket, h.object)

	return h
}

func (h *deleteObjectHandler) Process () error {
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

