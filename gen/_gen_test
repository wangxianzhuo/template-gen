package gen

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		wr                 io.Writer
		templateFileName   string
		templateConfigElem templateConfigElem
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				wr:               os.Stdout,
				templateFileName: "../sample/templateService",
				templateConfigElem: templateConfigElem{
					PackageName:  "com.sykj.river.master.service",
					ClassRoot:    "com.sykj.river.master",
					ClassName:    "River",
					ObjectName:   "river",
					GenerateDate: "2020-05-21T14:30:50+08:00",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args)
			if err := generate(tt.args.wr, tt.args.templateFileName, tt.args.templateConfigElem); (err != nil) != tt.wantErr {
				t.Errorf("generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseTemplate(t *testing.T) {
	type args struct {
		templateConfigData TemplateConfig
	}
	tests := []struct {
		name    string
		args    args
		wantWr  string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				templateConfigData: TemplateConfig{
					{
						PackageName:      "com.sykj.river.master.service",
						ClassRoot:        "com.sykj.river.master",
						ClassName:        "River",
						ObjectName:       "river",
						GenerateDate:     "2020-05-21T14:30:50+08:00",
						TemplateFilePath: "../sample/templateService",
					},
					{
						PackageName:      "com.sykj.river.master.service",
						ClassRoot:        "com.sykj.river.master",
						ClassName:        "River",
						ObjectName:       "river",
						GenerateDate:     "2020-05-21T14:30:50+08:00",
						TemplateFilePath: "../sample/templateServiceImpl",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &bytes.Buffer{}
			if err := ParseTemplate(wr, tt.args.templateConfigData); (err != nil) != tt.wantErr {
				t.Errorf("ParseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if gotWr := wr.String(); gotWr != tt.wantWr {
			// 	t.Errorf("ParseTemplate() = %v, want %v", gotWr, tt.wantWr)
			// }
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
		// TODO: Add test cases.
		{
			args: args{
				templateConfig: TemplateConfig{
					{
						PackageName:      "com.sykj.river.master.service",
						ClassRoot:        "com.sykj.river.master",
						ClassName:        "River",
						ObjectName:       "river",
						GenerateDate:     "2020-05-21T14:30:50+08:00",
						TemplateFilePath: "../sample/templateService",
						OutputFile:       "../sample/RiverService.Java",
					},
					{
						PackageName:      "com.sykj.river.master.service",
						ClassRoot:        "com.sykj.river.master",
						ClassName:        "River",
						ObjectName:       "river",
						GenerateDate:     "2020-05-21T14:30:50+08:00",
						TemplateFilePath: "../sample/templateServiceImpl",
						OutputFile:       "../sample/RiverServiceImpl.Java",
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
