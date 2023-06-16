# 替换空格和空行
sed -i '/^\s*$/d;s/[[:space:]]\+//g'  /tmp/judge11-res.log
sed -i '/^\s*$/d;s/[[:space:]]\+//g'  /data/workspace/myshixun/step/step11/judge_ans/ans.log

