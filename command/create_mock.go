package main

import (
	"fmt"
	"io/ioutil"

	"github.com/garyburd/redigo/redis"
	"gopkg.in/vmihailenco/msgpack.v2"
	"gopkg.in/yaml.v2"
)

type Config struct {
	IndexConfig IndexConfig `yaml:"index"`
}

type IndexConfig struct {
	BannerRegular   Index `yaml:"banner"`
	BannerRectangle Index `yaml:"banner_rect"`
	Native          Index `yaml:"native"`
	Video           Index `yaml:"video"`
}

type Index []Ad

type Ad struct {
	CampaignID string                 `yaml:"campaign_id"`
	CreativeID string                 `yaml:"creative_id"`
	Price      float64                `yaml:"price"`
	AdID       string                 `yaml:"ad_id"`
	NURL       string                 `yaml:"nurl"`
	IURL       string                 `yaml:"iurl"`
	AdM        string                 `yaml:"adm"`
	Adomain    map[string]interface{} `yaml:"adomain"`
	PeCPM      float64                `yaml:"pecpm"`
}

func main() {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	buf, err := ioutil.ReadFile("data.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.IndexConfig.BannerRegular != nil {
		b, err := msgpack.Marshal(config.IndexConfig.BannerRegular)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Do("SET", "index:banner", b)
	}

	if config.IndexConfig.BannerRectangle != nil {
		b, err := msgpack.Marshal(config.IndexConfig.BannerRectangle)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Do("SET", "index:banner_rect", b)
	}

	if config.IndexConfig.Native != nil {
		b, err := msgpack.Marshal(config.IndexConfig.Native)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Do("SET", "index:native", b)
	}

	if config.IndexConfig.Video != nil {
		b, err := msgpack.Marshal(config.IndexConfig.Video)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Do("SET", "index:native", b)
	}
}
