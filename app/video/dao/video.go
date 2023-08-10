package dao

import (
	"context"
	pb "dzug/protos/video"
)

func GetVideoFeed(ctx context.Context, req *pb.GetVideoFeedReq) (*pb.GetVideoFeedResp, error) {

	return nil, nil

}

/*// MGetVideoByTime 通过指定latestTime和count，从DAO层获取视频基本信息，并查出当前用户是否点赞，组装后返回
func (s *MGetVideoByTimeService) MGetVideoByTime(req *videoproto.GetVideoListByTimeReq) ([]*videoproto.VideoInfo, int64, error) {
	span := Tracer.StartSpan("feed")
	defer span.Finish()
	s.ctx = opentracing.ContextWithSpan(s.ctx, span)
	videoModels, nextTime, err := dal.MGetVideoByTime(s.ctx, time.Unix(req.LatestTime, 0), req.Count)
	// 只能得到视频id，uid，title，play_url,cover_url,created_time
	if err != nil {
		return nil, 0, err
	}
	videos := pack.Videos(videoModels) // 类型转换：视频id、base_info、点赞数、评论数已经得到，还需要判断是否点赞

	appUserID := req.AppUserId
	// 没有登录，直接返回不再查询是否点赞
	if appUserID < 0 {
		return videos, nextTime, nil
	}
	isLikeKeyExist, err := redis.IsLikeKeyExist(appUserID)
	if err != nil {
		klog.Error(err)
	}
	if isLikeKeyExist == false {
		// 如果redis没有appUserID的记录，则去mysql查询一次点赞列表进行缓存
		likeList, err := dal.MGetLikeList(s.ctx, appUserID)
		if err != nil {
			return nil, 0, err
		}
		if err := redis.AddLikeList(appUserID, likeList); err != nil {
			klog.Error(err)
		}
	}
	for i := 0; i < len(videos); i++ {
		isFavorite, err := redis.GetIsLikeById(appUserID, videos[i].VideoId)
		if err != nil {
			isFavorite, _ = dal.IsFavorite(s.ctx, videos[i].VideoId, appUserID)
		}
		videos[i].IsFavorite = isFavorite
	}
	return videos, nextTime, nil
}
*/