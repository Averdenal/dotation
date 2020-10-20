package logic

import (
	"github.com/Averdenal/Dotation/Models"
	"reflect"
	"testing"
)

func TestCreatedOrdinateur(t *testing.T) {
	type args struct {
		nameOrdinateur  string
		codeExpress     string
		os              string
		nbserial        string
		modelOrdinateur string
		tarif           string
	}
	tests := []struct {
		name    string
		args    args
		want    Models.Ordinateur
		wantErr bool
	}{
		{
			name: "Valide Test",
			args: args{
				nameOrdinateur:  "bgsa-a-154-prt",
				codeExpress:     "548745852",
				os:              "windows 10",
				nbserial:        "587445df",
				modelOrdinateur: "7450",
				tarif:           "1125.50",
			},
			want: Models.Ordinateur{
				Name:        "BGSA-A-154-PRT",
				CodeExpress: "548745852",
				Os:          "WINDOWS 10",
				Cat: Models.Cat{
					Name:   "Ordinateur",
					Models: "7450",
				},
				Tarif:    1125.5,
				NbSerial: "587445df",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreatedOrdinateur(tt.args.nameOrdinateur, tt.args.codeExpress, tt.args.os, tt.args.nbserial, tt.args.modelOrdinateur, tt.args.tarif)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatedOrdinateur() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("CreatedOrdinateur() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.Os, tt.want.Os) {
				t.Errorf("CreatedOrdinateur() got = %v, want %v", got.Os, tt.want.Os)
			}
			if !reflect.DeepEqual(got.Cat.Name, tt.want.Cat.Name) {
				t.Errorf("CreatedOrdinateur() got = %v, want %v", got.Cat.Name, tt.want.Cat.Name)
			}
		})
	}
}
