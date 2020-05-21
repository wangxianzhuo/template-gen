package gen

import (
	"bytes"
	"log"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		templateFileName string
		data             map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantWr  string
		wantErr bool
	}{
		{
			args: args{
				templateFileName: "../sample/templateService.gotmpl",
				data: map[string]string{
					"packageName":  "com.sykj.river.master.service",
					"classRoot":    "com.sykj.river.master",
					"className":    "River",
					"objectName":   "river",
					"generateDate": "2020-05-21T14:30:50+08:00",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &bytes.Buffer{}
			if err := generate(wr, tt.args.templateFileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if gotWr := wr.String(); gotWr != tt.wantWr {
			// 	t.Errorf("generate() = %v, want %v", gotWr, tt.wantWr)
			// }
			log.Println(wr.String())
		})
	}
}

func TestParseTemplate(t *testing.T) {
	type args struct {
		templateConfig TemplateConfig
	}
	tests := []struct {
		name    string
		args    args
		wantWr  string
		wantErr bool
	}{
		{
			args: args{
				templateConfig: TemplateConfig{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &bytes.Buffer{}
			if err := ParseTemplate(wr, tt.args.templateConfig); (err != nil) != tt.wantErr {
				t.Errorf("ParseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if gotWr := wr.String(); gotWr != tt.wantWr {
			// 	t.Errorf("ParseTemplate() = %v, want %v", gotWr, tt.wantWr)
			// }
			log.Println(wr.String())
		})
	}
}

func TestParseTemplateWriteToFile(t *testing.T) {
	type args struct {
		templateConfig TemplateConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				templateConfig: TemplateConfig{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseTemplateWriteToFile(tt.args.templateConfig); (err != nil) != tt.wantErr {
				t.Errorf("ParseTemplateWriteToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
