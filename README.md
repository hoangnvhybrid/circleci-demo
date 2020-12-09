# circleci-demo

Trong cấu hình Circle CI ta sẽ chỉ định các docker image (hoặc machine) sẽ sử dụng và các job, trong các job lại có các step, trong các step là cụ thể các command
Ngoài ra còn có cấu hình filter giúp ta linh hoạt điều chỉnh sao cho chỉ run các job khi có merge/push vào 1 số branch nhất định vân vân
 
Mô tả quá trình run 1 job trên Circle CI:

1. Developer chỉ cần push hoặc merge vào 1 branch, Circle CI tự động biết event đó và khởi động lên job đã được cài đặt tương ứng.
2. Ban đầu Circle CI pull docker image về và run lên trên môi trường cloud của nó
3. Tiếp theo nó chạy các step đã được cài đặt trong docker container, thông thường step đầu tiên luôn là checkout tức là git checkout lấy source về (mặc định lưu trong thư mục ~/project)
4. Các step tiếp theo được chạy tùy vào độ sáng tạo của bạn, ví dụ job để build thì thường là npm install rồi npm run abcxyz hay job để deploy thì có thể là aws s3 sync hay serverless deploy...
5. Sau khi tất cả các step đã chạy xong, job kết thúc. Nếu exit code của job là error thì mặc định ta sẽ nhận được mail thông báo failed nữa.
