package main
import "fmt"
func main(){
//int
	o := make(map[int]int)
	o[1] =23
	o[3] =29
	o[2] =21
	fmt.Println("o=",o)
//key-int,val-int
	p :=make(map[int]int)

	p[5] =123
	p[1] =923
	p[2] =38
	p[4] =838847
	p[6]  =38738
	p[3] =98383
	fmt.Println("p",p)
//key-float ,val-int
	z :=make(map[float32]int)
	z[1.1] =9822
	z[1.3] =3884
	z[1.2] =239838
	z[3.1] =28
	z[2.3]=298
	z[3.1]=2983
	z[2.1]=12334
	z[3.3]=983239
	z[2.2]=7237 
	val3 ,ok:=z[3.3]
	fmt.Println("float z =",z,"val3",val3,"ok",ok)
	val3 ,ok = z[3.4]
	if ok{
		fmt.Println("val3 in key [3.4]=",val3 ,"ok=",ok)
	}else{
		fmt.Println("key and value is not found")
	}

	//key-string,val-int

	s :=make(map[string]int)
	s["one"]=89932
	s["three"]=1234
	s["two"]=93839

	fmt.Println("string s=",s)
//key - int,value-float32
	r :=make(map[int]float32)
	r [1]=98238
	r[2]=7934
	r[4]=0.1232
	r[3]=74
	val1 ,ok:=r[4]
	fmt.Println("val1=",val1,"ok=",ok)
//key-int.,value-string
	r1 :=make(map[int]string)
	r1 [1]="hello g"
	r1[2]="my name"
	r1[4]="thakur"
	r1[3]="prashant"
    val2 ,ok:=r1[4]
	fmt.Println("val2=",val2,"ok=",ok,"key -int . value - string",r1)



}