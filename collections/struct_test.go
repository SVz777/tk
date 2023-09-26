package collections

import (
	"reflect"
	"testing"
)

func TestGetValueWithFieldPath(t *testing.T) {
	type args struct {
		data any
		path []string
	}
	type DataOneLayer struct {
		A int
	}

	type Multi1 struct {
		A int
	}
	type Multi2 struct {
		A any
	}
	type Multi3 struct {
		A any
	}
	type DataMultiLayer struct {
		A any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "path error",
			args: args{
				data: DataOneLayer{
					A: 1,
				},
				path: []string{"A.A"},
			},
			want:    1,
			wantErr: true,
		},
		{
			name: "one layer get A",
			args: args{
				data: DataOneLayer{
					A: 1,
				},
				path: []string{"A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "*one layer get A",
			args: args{
				data: &DataOneLayer{
					A: 1,
				},
				path: []string{"A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 1",
			args: args{
				data: DataMultiLayer{
					A: Multi1{A: 1},
				},
				path: []string{"A", "A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 1 *",
			args: args{
				data: DataMultiLayer{
					A: &Multi1{A: 1},
				},
				path: []string{"A", "A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 2",
			args: args{
				data: DataMultiLayer{
					A: Multi2{
						A: Multi1{
							A: 1,
						},
					},
				},
				path: []string{"A", "A", "A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 2 *-",
			args: args{
				data: DataMultiLayer{
					A: &Multi2{
						A: Multi1{
							A: 1,
						},
					},
				},
				path: []string{"A", "A", "A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 2 -*",
			args: args{
				data: DataMultiLayer{
					A: Multi2{
						A: &Multi1{
							A: 1,
						},
					},
				},
				path: []string{"A", "A", "A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 2 **",
			args: args{
				data: DataMultiLayer{
					A: &Multi2{
						A: &Multi1{
							A: 1,
						},
					},
				},
				path: []string{"A", "A", "A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 3",
			args: args{
				data: DataMultiLayer{
					A: Multi3{
						A: Multi2{
							A: Multi1{
								A: 1,
							},
						},
					},
				},
				path: []string{"A", "A", "A", "A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 3 ***",
			args: args{
				data: DataMultiLayer{
					A: &Multi3{
						A: &Multi2{
							A: &Multi1{
								A: 1,
							},
						},
					},
				},
				path: []string{"A", "A", "A", "A"},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "multi layer 3 *-*",
			args: args{
				data: DataMultiLayer{
					A: &Multi3{
						A: Multi2{
							A: &Multi1{
								A: 1,
							},
						},
					},
				},
				path: []string{"A", "A", "A", "A"},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := GetValueWithFieldPath(tt.args.data, tt.args.path...)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetValueWithFieldPath() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err == nil && !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetValueWithFieldPath() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestSetValueWithFieldPath(t *testing.T) {
	type args struct {
		data  any
		value any
		path  []string
	}
	type DataOneLayer struct {
		A any
		I int
	}

	type DataMultiLayer1 struct {
		A any
	}
	type DataMultiLayer2 struct {
		A any
	}
	type DataMultiLayer3 struct {
		A any
	}
	tests := []struct {
		name      string
		args      args
		wantValue any
		wantErr   bool
	}{
		{
			name: "one layer kind not equal",
			args: args{
				data:  &DataOneLayer{},
				value: "A",
				path:  []string{"I"},
			},
			wantValue: "A",
			wantErr:   true,
		},
		{
			name: "one layer set A int",
			args: args{
				data: &DataOneLayer{
					A: nil,
					I: 0,
				},
				value: 2,
				path:  []string{"A"},
			},
			wantValue: 2,
			wantErr:   false,
		},
		{
			name: "one layer set A interface",
			args: args{
				data:  &DataOneLayer{},
				value: "a",
				path:  []string{"A"},
			},
			wantValue: "a",
			wantErr:   false,
		},
		{
			name: "multi layer 1 set A int",
			args: args{
				data: DataMultiLayer1{
					A: &DataOneLayer{
						A: "e",
						I: 0,
					},
				},
				value: 1,
				path:  []string{"A", "I"},
			},
			wantValue: 1,
			wantErr:   false,
		},
		{
			name: "multi layer 1 set A interface ",
			args: args{
				data: DataMultiLayer1{
					A: &DataOneLayer{
						A: "e",
						I: 0,
					},
				},
				value: "s",
				path:  []string{"A", "A"},
			},
			wantValue: "s",
			wantErr:   false,
		},
		{
			name: "multi layer 1 can't set ",
			args: args{
				data: DataMultiLayer1{
					A: DataOneLayer{
						A: "e",
						I: 0,
					},
				},
				value: "s",
				path:  []string{"A", "A"},
			},
			wantValue: "s",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := SetValueWithFieldPath(
					tt.args.data,
					tt.args.value,
					tt.args.path...,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("SetValueWithFieldPath() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err != nil {
					return
				}
				if v, _ := GetValueWithFieldPath(tt.args.data, tt.args.path...); v != tt.wantValue {
					t.Errorf("SetValueWithFieldPath() value = %+v, wantValue %+v", v, tt.wantValue)
				}
			},
		)
	}
}
