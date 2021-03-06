// Copyright 2016 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/i18n/bundle"
	"github.com/spf13/hugo/source"
	"github.com/spf13/hugo/tpl"
	jww "github.com/spf13/jwalterweatherman"
)

func loadI18n(sources []source.Input) error {
	jww.DEBUG.Printf("Load I18n from %q", sources)

	i18nBundle := bundle.New()

	for _, currentSource := range sources {
		for _, r := range currentSource.Files() {
			err := i18nBundle.ParseTranslationFileBytes(r.LogicalName(), r.Bytes())
			if err != nil {
				return fmt.Errorf("Failed to load translations in file %q: %s", r.LogicalName(), err)
			}
		}
	}

	tpl.SetI18nTfuncs(i18nBundle)

	return nil

}
