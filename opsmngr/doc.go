// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package opsmngr provides a client for using the MongoDB Ops Manager and Cloud Manager API.

Usage

	import "go.mongodb.org/ops-manager/opsmngr"

Construct a new Ops Manager client, then use the various services on the client to
access different parts of the Ops Manager API. For example:

	client := opsmngr.NewClient(nil)

The services of a client divide the API into logical chunks and correspond to
the structure of the Ops Manager API documentation at
https://docs.opsmanager.mongodb.com/v4.2/reference/api/.

NOTE: Using the https://godoc.org/context package, one can easily
pass cancellation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then context.Background()
can be used as a starting point.

Authentication

The ops-manager library does not directly handle authentication. Instead, when
creating a new client, pass an http.Client that can handle Digest Access authentication for
you. The easiest way to do this is using the https://github.com/mongodb-forks/digest
library, but you can always use any other library that provides an `http.Client`.
If you have a private and public API token pair, you can use it with the digest library using:

	import (
		"context"
		"log"

		"github.com/mongodb-forks/digest"
		"go.mongodb.org/ops-manager/opsmngr"
	)

	func main() {
		t := digest.NewTransport("your public key", "your private key")
		tc, err := t.Client()
		if err != nil {
			log.Fatalf(err.Error())
		}

		// Note: If no Base URL is set the client is set to work with Cloud Manager by default
		clientops := opsmngr.SetBaseURL("https://opsmanagerurl/" + opsmngr.APIPublicV1Path)
		client, err := opsmngr.New(tc, clientops)
		orgs, _, err := client.Organizations.List(context.Background(), nil)
	}

Note that when using an authenticated Client, all calls made by the client will
include the specified tokens. Therefore, authenticated clients should
almost never be shared between different users.

To get your API Keys please refer to our docs for:

- Ops Manager: https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/,

- Cloud Manager: https://docs.cloudmanager.mongodb.com/tutorial/manage-programmatic-api-keys/.
*/
package opsmngr
