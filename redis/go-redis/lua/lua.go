package lua

var ZsetADDMember = `
-- param : key expire score1 member1 score2 member2 ...

-- unpack 函数(用于解包数据)
local unpack = _G.unpack or table.unpack
-- pack 函数(用于打包数据)
local pack = function(...)
    return { n = select("#", ...), ... }
end
local key = KEYS[1]
local expire = ARGV[1]
local members = pack(unpack(ARGV, 2))

-- 检查排行榜是否存在
local exists = redis.call("EXISTS", key)

redis.call("ZADD", key, unpack(members))

if exists == 0 then
    if tonumber(expire) > 1 then
        redis.call("EXPIRE", key, expire)
    end
end
return 1
`

// sptGetMembers 获取成员排行以及积分信息
var sptGetMembers = `
local name   = KEYS[1]
local member = ARGV[1]

local score = redis.call('ZSCORE', name, member)
local rank = redis.call('ZREVRANK', name, member)

return {tonumber(score), tonumber(rank)}
`

var sptIncr = `
-- param : key expire score member

local key = KEYS[1]
local expire = ARGV[1]
local score = ARGV[2]
local member = ARGV[3]

local exists = redis.call("EXISTS", key)

redis.call("ZINCRBY", key, score, member)

if exists == 0 then
    if tonumber(expire) > 1 then
        redis.call("EXPIRE", key, expire)
    end
end

return 1
`
