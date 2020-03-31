package template

import "testing"

type Person struct {
	Name string
	Age int
	Likes string
}

var Levin = Person{
	Name: "Levin",
	Age: 32,
	Likes: "Computers",
}

func TestParse(t *testing.T) {
	type args struct {
		goTemplate string
		data       interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should generate from template",
			args:args{
				goTemplate: "My Name is {{.Name}}, I'm {{.Age}} and I like {{.Likes}}",
				data: Levin,
			},
			want: "My Name is Levin, I'm 32 and I like Computers",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.goTemplate, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
