
 k exec redis-worker-5c8b468cf7-5sk9r  -c worker  -- redis-cli  -p 6379  set fop 304

 k exec redis-worker-5c8b468cf7-5sk9r  -it -- redis-cli info memory

 k exec redis-worker-5c8b468cf7-5sk9r  -it -- redis-cli info server

 k exec redis-worker-5c8b468cf7-5sk9r  -it -- redis-cli info keyspace

