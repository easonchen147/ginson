package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"ginson/app/model"
	"time"
)

type WxMiniCache struct {
	*BaseCache
}

var wxMiniCache = &WxMiniCache{BaseCache: baseCache}

func GetWxMiniCache() *WxMiniCache {
	return wxMiniCache
}

func (c *WxMiniCache) getOpenIdKey(openId string) string {
	return fmt.Sprintf("openId:%s", openId)
}

func (c *WxMiniCache) AddSessionKey(ctx context.Context, openId, sessionKey string) error {
	return c.redis().Set(ctx, c.getOpenIdKey(openId), sessionKey, time.Minute*30).Err()
}

func (c *WxMiniCache) GetSessionKey(ctx context.Context, openId string) (string, error) {
	return c.redis().Get(ctx, fmt.Sprintf("openId:%s", openId)).Result()
}

func (c *WxMiniCache) getUserInfoKey(openId string) string {
	return fmt.Sprintf("userInfo:%s", openId)
}

func (c *WxMiniCache) AddUserInfo(ctx context.Context, openId string, userInfo *model.WxMiniGetUserInfoResp) error {
	jsonBytes, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	return c.redis().Set(ctx, c.getUserInfoKey(openId), string(jsonBytes), time.Hour).Err()
}

func (c *WxMiniCache) GetUserInfo(ctx context.Context, openId string) (*model.WxMiniGetUserInfoResp, error) {
	jsonStr, err := c.redis().Get(ctx, c.getUserInfoKey(openId)).Result()
	if err != nil {
		return nil, err
	}

	result := &model.WxMiniGetUserInfoResp{}
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
