// Copyright 2020 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "agones.dev/agones/pkg/apis/agones/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GameServerLister helps list GameServers.
// All objects returned here must be treated as read-only.
type GameServerLister interface {
	// List lists all GameServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GameServer, err error)
	// GameServers returns an object that can list and get GameServers.
	GameServers(namespace string) GameServerNamespaceLister
	GameServerListerExpansion
}

// gameServerLister implements the GameServerLister interface.
type gameServerLister struct {
	indexer cache.Indexer
}

// NewGameServerLister returns a new GameServerLister.
func NewGameServerLister(indexer cache.Indexer) GameServerLister {
	return &gameServerLister{indexer: indexer}
}

// List lists all GameServers in the indexer.
func (s *gameServerLister) List(selector labels.Selector) (ret []*v1.GameServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GameServer))
	})
	return ret, err
}

// GameServers returns an object that can list and get GameServers.
func (s *gameServerLister) GameServers(namespace string) GameServerNamespaceLister {
	return gameServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GameServerNamespaceLister helps list and get GameServers.
// All objects returned here must be treated as read-only.
type GameServerNamespaceLister interface {
	// List lists all GameServers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GameServer, err error)
	// Get retrieves the GameServer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GameServer, error)
	GameServerNamespaceListerExpansion
}

// gameServerNamespaceLister implements the GameServerNamespaceLister
// interface.
type gameServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GameServers in the indexer for a given namespace.
func (s gameServerNamespaceLister) List(selector labels.Selector) (ret []*v1.GameServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GameServer))
	})
	return ret, err
}

// Get retrieves the GameServer from the indexer for a given namespace and name.
func (s gameServerNamespaceLister) Get(name string) (*v1.GameServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("gameserver"), name)
	}
	return obj.(*v1.GameServer), nil
}
