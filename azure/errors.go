package azure

import "errors"

var PlatformStorageError = errors.New("When using a platform image, the 'storage_service_name' parameter is required")
