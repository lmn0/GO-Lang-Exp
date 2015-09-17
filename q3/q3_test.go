package main
import ("testing";"time")

func TestSleep(t *testing.T){

	start:=time.Now();
	Sleep(3);
	watch:=time.Since(start);

	if int64(watch)/1000000000!=int64(time.Second * time.Duration(3))/1000000000{
		t.Error("Oops")
	}
}
