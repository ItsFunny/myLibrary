/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-19 13:24
# @File : randutil.go
# @Description :
# @Attention :
*/
package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestString(t *testing.T) {
	s, e := String(4, "qwertyui")
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println(s)
}

func TestStringRange(t *testing.T) {
	type args struct {
		min     int
		max     int
		charset string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringRange(tt.args.min, tt.args.max, tt.args.charset)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlphaStringRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AlphaStringRange(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("AlphaStringRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AlphaStringRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlphaString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AlphaString(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("AlphaString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AlphaString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChoiceString(t *testing.T) {
	type args struct {
		choices []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChoiceString(tt.args.choices)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChoiceString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ChoiceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
