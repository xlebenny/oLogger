package ologger

import (
	"testing"
)

func TestLog(t *testing.T) {
	type innerStruct struct {
		fieldC string `oLog:"0"`
	}

	type args struct {
		logLevel     int
		indentString string
		obj          interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestLogExportVariable",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					FieldA string `oLog:"0"`
				}{
					FieldA: "valueA",
				},
			},
			want: `{"":{"FieldA":"valueA"}}`,
		},
		{
			name: "TestLogUnexportVariable",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					fieldA string `oLog:"0"`
				}{
					fieldA: "valueA",
				},
			},
			want: `{"":{"fieldA":"valueA"}}`,
		},
		{
			name: "TestLogTwoVariable",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					fieldA string `oLog:"0"`
					fieldB int    `oLog:"0"`
				}{
					fieldA: "valueA",
					fieldB: 123,
				},
			},
			want: `{"":{"fieldA":"valueA","fieldB":123}}`,
		},
		{
			name: "TestLogInnerVariable",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					fieldA string      `oLog:"0"`
					fieldB innerStruct `oLog:"0"`
				}{
					fieldA: "valueA",
					fieldB: innerStruct{
						fieldC: "valueC",
					},
				},
			},
			want: `{"":{"fieldA":"valueA","fieldB":{"fieldC":"valueC"}}}`,
		},
		{
			name: "TestLogLevel < oLog",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					fieldA string `oLog:"5"`
					fieldB int    `oLog:"0"`
				}{
					fieldA: "valueA",
					fieldB: 123,
				},
			},
			want: `{"":{"fieldB":123}}`,
		},
		{
			name: "TestLogLevel == oLog",
			args: args{
				logLevel:     5,
				indentString: "",
				obj: struct {
					fieldA string `oLog:"5"`
					fieldB int    `oLog:"0"`
				}{
					fieldA: "valueA",
					fieldB: 123,
				},
			},
			want: `{"":{"fieldA":"valueA","fieldB":123}}`,
		},
		{
			name: "TestLogLevel > oLog",
			args: args{
				logLevel:     6,
				indentString: "",
				obj: struct {
					fieldA string `oLog:"5"`
					fieldB int    `oLog:"0"`
				}{
					fieldA: "valueA",
					fieldB: 123,
				},
			},
			want: `{"":{"fieldA":"valueA","fieldB":123}}`,
		},
		{
			name: "TestLogLevelIsAlphabet",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					fieldA string `oLog:"A"`
					fieldB int    `oLog:"0"`
				}{
					fieldA: "valueA",
					fieldB: 123,
				},
			},
			want: `{"":{"fieldB":123}}`,
		},
		{
			name: "TestLogLevelIsEmpty",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					fieldA string `oLog:""`
					fieldB int    `oLog:"0"`
				}{
					fieldA: "valueA",
					fieldB: 123,
				},
			},
			want: `{"":{"fieldB":123}}`,
		},
		{
			name: "TestNoOLogTag",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					fieldA string ``
					fieldB int    `oLog:"0"`
				}{
					fieldA: "valueA",
					fieldB: 123,
				},
			},
			want: `{"":{"fieldB":123}}`,
		},
		{
			name: "TestTagCaseInsensitive",
			args: args{
				logLevel:     4,
				indentString: "",
				obj: struct {
					fieldA string `OlOg:"0"`
					fieldB int    `oLoG:"0"`
				}{
					fieldA: "valueA",
					fieldB: 123,
				},
			},
			want: `{"":{}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Log(tt.args.logLevel, tt.args.indentString, tt.args.obj); got != tt.want {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}
