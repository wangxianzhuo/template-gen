package gen

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	type args struct {
		config []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult TemplateConfig
	}{
		{
			args: args{
				config: []byte(`[
					{
						"templateFile": "../sample/templateService.gotmpl",
						"outputFile": "../sample/RiverService.Java",
						"renderData": {
							"packageName": "com.sykj.river.master.service",
							"classRoot": "com.sykj.river.master",
							"className": "River",
							"objectName": "river",
							"generateDate": "2020-05-21T14:30:50+08:00"
						}
					}
				]`),
			},
			wantResult: TemplateConfig{
				{
					TemplateFile: "../sample/templateService.gotmpl",
					OutputFile:   "../sample/RiverService.Java",
					RenderData: map[string]string{
						"packageName":  "com.sykj.river.master.service",
						"classRoot":    "com.sykj.river.master",
						"className":    "River",
						"objectName":   "river",
						"generateDate": "2020-05-21T14:30:50+08:00",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ParseConfig(tt.args.config); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ParseConfig() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestParseConfigFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name       string
		args       args
		wantResult TemplateConfig
	}{
		{
			args: args{
				fileName: "../sample/config_sample.json",
			},
			wantResult: TemplateConfig{
				{
					TemplateFile: "../sample/templateService.gotmpl",
					OutputFile:   "../sample/RiverService.Java",
					RenderData: map[string]string{
						"packageName":  "com.sykj.river.master.service",
						"classRoot":    "com.sykj.river.master",
						"className":    "River",
						"objectName":   "river",
						"generateDate": "2020-05-21T14:30:50+08:00",
					},
				},
				{
					TemplateFile: "../sample/templateServiceImpl.gotmpl",
					OutputFile:   "../sample/RiverServiceImpl.Java",
					RenderData: map[string]string{
						"packageName":  "com.sykj.river.master.service",
						"classRoot":    "com.sykj.river.master",
						"className":    "River",
						"objectName":   "river",
						"generateDate": "2020-05-21T14:30:50+08:00",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ParseConfigFile(tt.args.fileName); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ParseConfigFile() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
