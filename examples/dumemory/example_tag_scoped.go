/*
 * Copyright 2026 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing
 * permissions
 * and limitations under the License.
 */

package dumemoryexamples

import (
	"context"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/dumemory/api"
)

func RetainWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AppID: "your-app-id"}

	item := *dumemory.NewMemoryItem("User likes concise technical explanations.")
	item.Tags = []string{"topic:preference"}
	req := dumemory.NewRetainRequest([]dumemory.MemoryItem{item})
	req.DocumentTags = []string{"source:example"}

	response, err := client.RetainWithScope(context.Background(), "your-bank-id", scope, *req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func RecallWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AppID: "your-app-id"}

	req := dumemory.NewRecallRequest("What response style does the user prefer?")
	req.Tags = []string{"topic:preference"}
	response, err := client.RecallWithScope(context.Background(), "your-bank-id", scope, *req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func ReflectWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AgentID: "your-agent-id"}

	req := dumemory.NewReflectRequest("Summarize the user's current working preferences.")
	req.Tags = []string{"topic:preference"}
	response, err := client.ReflectWithScope(context.Background(), "your-bank-id", scope, *req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func ListTagsWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id"}

	response, err := client.ListTagsWithScope(context.Background(), "your-bank-id", scope, dumemory.ListTagsOptions{
		Source: "memories",
		Limit:  20,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func ListMemoriesWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AppID: "your-app-id"}

	response, err := client.ListMemoriesWithScope(context.Background(), "your-bank-id", scope, dumemory.ListMemoriesOptions{
		Q:          "preference",
		State:      "active",
		DocumentID: "your-document-id",
		EntityID:   "your-entity-id",
		Tags:       []string{"topic:preference"},
		Limit:      20,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func ListDocumentsWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", RunID: "your-run-id"}

	response, err := client.ListDocumentsWithScope(context.Background(), "your-bank-id", scope, dumemory.ListDocumentsOptions{
		Tags:  []string{"source:example"},
		Limit: 20,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func UpdateDocumentTagsWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AppID: "your-app-id"}

	req := dumemory.NewUpdateDocumentRequest()
	req.Tags = []string{"source:example", "topic:preference"}
	response, err := client.UpdateDocumentTagsWithScope(context.Background(), "your-bank-id", "your-document-id", scope, *req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func CreateDirectiveWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AgentID: "your-agent-id"}

	req := dumemory.NewCreateDirectiveRequest("scoped-tone", "Use a concise and practical tone.")
	req.Tags = []string{"topic:style"}
	response, err := client.CreateDirectiveWithScope(context.Background(), "your-bank-id", scope, *req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func ListDirectivesWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AgentID: "your-agent-id"}

	response, err := client.ListDirectivesWithScope(context.Background(), "your-bank-id", scope, dumemory.ListDirectivesOptions{
		Tags:       []string{"topic:style"},
		ActiveOnly: true,
		Limit:      20,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func UpdateDirectiveWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AgentID: "your-agent-id"}

	req := dumemory.NewUpdateDirectiveRequest()
	req.Tags = []string{"topic:style", "state:active"}
	response, err := client.UpdateDirectiveWithScope(context.Background(), "your-bank-id", "your-directive-id", scope, *req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func CreateMentalModelWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AppID: "your-app-id"}

	req := dumemory.NewCreateMentalModelRequest("User Preference Model", "What does this user prefer while working?")
	req.Tags = []string{"topic:preference"}
	response, err := client.CreateMentalModelWithScope(context.Background(), "your-bank-id", scope, *req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func ListMentalModelsWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AppID: "your-app-id"}

	response, err := client.ListMentalModelsWithScope(context.Background(), "your-bank-id", scope, dumemory.ListMentalModelsOptions{
		Tags:  []string{"topic:preference"},
		Limit: 20,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func UpdateMentalModelWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id", AppID: "your-app-id"}

	req := dumemory.NewUpdateMentalModelRequest()
	req.Tags = []string{"topic:preference", "state:active"}
	response, err := client.UpdateMentalModelWithScope(context.Background(), "your-bank-id", "your-model-id", scope, *req)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func GetMemoryWithScope() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := dumemory.New(endpoint, apiKey)
	scope := dumemory.EntityScope{UserID: "your-user-id"}

	response, err := client.GetMemoryWithScope(context.Background(), "your-bank-id", "your-memory-id", scope)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
