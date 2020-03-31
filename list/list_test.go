package list

import (
"testing"
)

func Test_Contains(t *testing.T) {
    type args struct {
        target interface{}
        values []interface{}
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "should be true if equal to one of",
            args: args{
                target: "red",
                values: []interface{}{"blue", "green", "red"},
            },
            want: true,
        },
        {
            name: "should be false if equal to none of",
            args: args{
                target: "blue",
                values: []interface{}{"green", "red"},
            },
            want: false,
        },
        {
            name: "should be true if equal to all of",
            args: args{
                target: "blue",
                values: []interface{}{"blue", "blue"},
            },
            want: true,
        },
        {
            name: "should be true if one of int",
            args: args{
                target: 5,
                values: []interface{}{1, 5},
            },
            want: true,
        },
        {
            name: "should be true if none of int",
            args: args{
                target: 5,
                values: []interface{}{1, 2},
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Contains(tt.args.target, tt.args.values...); got != tt.want {
                t.Errorf("contains() = %v, want %v", got, tt.want)
            }
        })
    }
}
