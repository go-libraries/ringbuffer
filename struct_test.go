package ringbuffer

import (
    "testing"
    "fmt"
)

type IName interface {
    SetName(name string)
    GetName() string
}
type Company struct{
    Name string    //可见的
    Age     int
}
type User struct {
    username string  //不可见的
}
//可见的
func (c *Company) SetName(name string){
    c.setName(name)
}
//不可见的
func (c *Company) setName(name string){
    c.Name = "name:"+name
}

func (c *Company) GetName() string {
    return c.Name
}
//可见的
func (c *User) SetName(name string){
    c.setName(name)
}
//不可见的
func (c *User) setName(name string){
    c.username = "name:"+name
}

func (c *User) GetName() string {
    return c.username
}


func TestCompany(t *testing.T) {
    c := &Company{
        Name:"北京蜜莱坞",
        Age:4,
    }
    c.SetName("北京蜜莱坞网络科技有限公司")

    u := &User{
        username:"李一文",
    }
    u.SetName("一文哥")

    showName(u)
    showName(c)
}

func showName(i IName) {
    fmt.Println(i.GetName())
}
