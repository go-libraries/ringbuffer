package ringbuffer

import (
    "testing"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "log"
    "go.mongodb.org/mongo-driver/bson"
    "fmt"
    "math"
    "runtime"
)

var client *mongo.Client

func init() {

    ctx, _ := context.WithTimeout(context.Background(), 300*time.Second)
    op := options.Client()
    op.SetMaxPoolSize(1000)
    c, err := mongo.Connect(ctx, op.ApplyURI(
        "mongodb://127.0.0.1.217:27017/test?w=majority&maxPoolSize=256",
    ))
    if err != nil { log.Fatal(err) }
    client = c
    client.StartSession()

}

func TestMongodb(t *testing.T) {
    maxProcess := runtime.NumCPU()
    fmt.Println(maxProcess)
    if maxProcess > 1 {
        maxProcess -= 1
    }
    runtime.GOMAXPROCS(maxProcess)
    client.StartSession()

    for i := 0; i < 5000; i++ {
        go func() {
            defer func() {
                // 如果程序异常, 停止当前定时任务,记录日志,重启任务
                if x := recover(); x != nil {
                    log.Println("do things :", x)
                }
            }()
            collection := client.Database("testing").Collection("taiex_399")
            ctx, exit := context.WithTimeout(context.Background(), 5*time.Second)
            defer exit()
            _, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": math.Round(10000)})
            if err != nil {
                fmt.Println(err)
                return
            }
            //id := res.InsertedID
            //fmt.Println(err, id)
        }()
    }

    //ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
    //cur, err := collection.Find(ctx, bson.M{"name":"pi"})
    //if err != nil { log.Fatal(err.Error()) }
    //defer cur.Close(ctx)
    //for cur.Next(ctx) {
    //    //var result bson.M
    //    fmt.Println(cur.Current)
    //    //err := cur.Decode(&result)
    //    //if err != nil { log.Fatal(err) }
    //    // do something with result....
    //}
    //if err := cur.Err(); err != nil {
    //    log.Fatal(err)
    //}

    select{}
}