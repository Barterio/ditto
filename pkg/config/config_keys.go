// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package config

const SERVER_1_ENDPOINT = "Server1.Endpoint"
const SERVER_1_ACCESS_KEY = "Server1.AccessKey"
const SERVER_1_SECRET_KEY = "Server1.SecretKey"

const SERVER_2_ENDPOINT = "Server2.Endpoint"
const SERVER_2_ACCESS_KEY = "Server2.AccessKey"
const SERVER_2_SECRET_KEY = "Server2.SecretKey"

const DEFAULT_OPTIONS_DEFAULT_SOURCE = "DefaultOptions.DefaultSource"
const DEFAULT_OPTIONS_THROW_IMMEDIATELY = "DefaultOptions.ThrowImmediately"

const LIST_DEFAULT_SOURCE = "ListOptions." + DEFAULT_OPTIONS_DEFAULT_SOURCE
const LIST_THROW_IMMEDIATELY = "ListOptions." + DEFAULT_OPTIONS_THROW_IMMEDIATELY
const LIST_MERGE = "ListOptions.Merge"

const PUT_DEFAULT_SOURCE = "PutOptions." + DEFAULT_OPTIONS_DEFAULT_SOURCE
const PUT_THROW_IMMEDIATELY = "PutOptions." + DEFAULT_OPTIONS_THROW_IMMEDIATELY
const PUT_CREATE_BUCKET_IF_NOT_EXIST = "PutOptions.CreateBucketIfNotExist"

const GET_OBJECT_DEFAULT_SOURCE = "GetObjectOptions." + DEFAULT_OPTIONS_DEFAULT_SOURCE
const GET_OBJECT_THROW_IMMEDIATELY = "GetObjectOptions." + DEFAULT_OPTIONS_THROW_IMMEDIATELY

const COPY_DEFAULT_SOURCE = "CopyOptions." + DEFAULT_OPTIONS_DEFAULT_SOURCE
const COPY_THROW_IMMEDIATELY = "CopyOptions." + DEFAULT_OPTIONS_THROW_IMMEDIATELY

const DELETE_DEFAULT_SOURCE = "DeleteOptions." + DEFAULT_OPTIONS_DEFAULT_SOURCE
const DELETE_THROW_IMMEDIATELY = "DeleteOptions." + DEFAULT_OPTIONS_THROW_IMMEDIATELY

// const ConfigKeys:= make(string, 20){"",""}
func GetKeysArray() []string {
	return []string{
		SERVER_1_ENDPOINT,
		SERVER_1_ACCESS_KEY,
		SERVER_1_SECRET_KEY,
		SERVER_2_ENDPOINT,
		SERVER_2_ACCESS_KEY,
		SERVER_2_SECRET_KEY,
		DEFAULT_OPTIONS_DEFAULT_SOURCE,
		DEFAULT_OPTIONS_THROW_IMMEDIATELY,
		LIST_DEFAULT_SOURCE,
		LIST_THROW_IMMEDIATELY,
		LIST_MERGE,
		PUT_DEFAULT_SOURCE,
		PUT_THROW_IMMEDIATELY,
		PUT_CREATE_BUCKET_IF_NOT_EXIST,
		GET_OBJECT_DEFAULT_SOURCE,
		GET_OBJECT_THROW_IMMEDIATELY,
		COPY_DEFAULT_SOURCE,
		COPY_THROW_IMMEDIATELY,
		DELETE_DEFAULT_SOURCE,
		DELETE_THROW_IMMEDIATELY,
	}
}
