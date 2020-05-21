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
		// TODO: Add test cases.
		{
			args: args{
				config: []byte(`[
					{
						"packageName": "com.sykj.river.master.service",
						"classRoot": "com.sykj.river.master",
						"className": "River",
						"objectName": "river",
						"generateDate": "2020-05-21T14:30:50+08:00",
						"templateFilePath": "../sample/templateService",
						"outputFile": "../sample/RiverService.Java",
					}
				]`),
			},
			wantResult: TemplateConfig{
				{
					PackageName:      "com.sykj.river.master.service",
					ClassRoot:        "com.sykj.river.master",
					ClassName:        "River",
					ObjectName:       "river",
					GenerateDate:     "2020-05-21T14:30:50+08:00",
					TemplateFilePath: "../sample/templateService",
					OutputFile:       "../sample/RiverService.Java",
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
		// TODO: Add test cases.
		{
			args: args{
				fileName: "../sample/config_sample.json",
			},
			wantResult: TemplateConfig{
				{
					PackageName:      "com.sykj.river.master.service",
					ClassRoot:        "com.sykj.river.master",
					ClassName:        "River",
					ObjectName:       "river",
					GenerateDate:     "2020-05-21T14:30:50+08:00",
					TemplateFilePath: "../sample/templateService",
					OutputFile:       "../sample/RiverService.Java",
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
