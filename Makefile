MONGO_SERVER ?= localhost:27010
ZOOKEEPER_SERVER ?= localhost:2181

build:
	go build -o main main.go

run:
	./main

test_float:
	go test -timeout 60s -run ^TestFloatSaveNumBasic$$ github.com/smiecj/go_basic/test -v -count=1

test_bench_float:
	go test -benchmem -benchtime=20s -bench='BenchmarkFloatSaveNum' -run=none github.com/smiecj/go_basic/test

test_slice_sub:
	go test -timeout 60s -run ^TestSubSlice$$ github.com/smiecj/go_basic/basic -v -count=1

test_slice_append:
	go test -timeout 60s -run ^TestSliceAppend$$ github.com/smiecj/go_basic/basic -v -count=1

test_bytes_to_string_bench:
	go test -benchmem -benchtime=10s -bench='BenchmarkForceConvertBytesToString' -run=none github.com/smiecj/go_basic/basic/string
	go test -benchmem -benchtime=10s -bench='BenchmarkConvertBytesToString' -run=none github.com/smiecj/go_basic/basic/string

test_string_to_bytes_bench:
	go test -benchmem -benchtime=10s -bench='BenchmarkForceConvertStringToBytes' -run=none github.com/smiecj/go_basic/basic/string
	go test -benchmem -benchtime=10s -bench='BenchmarkConvertStringToBytes' -run=none github.com/smiecj/go_basic/basic/string

test_map_get:
	go test -timeout 60s -run ^TestMapGet$$ github.com/smiecj/go_basic/basic -v -count=1

test_mutex_copy:
	go test -race -timeout 10s -run ^TestMutexCopy$$ github.com/smiecj/go_basic/sync_ -v -count=1

test_mutex_copy_simple:
	go test -race -timeout 10s -run ^TestMutexCopySimple$$ github.com/smiecj/go_basic/sync_ -v -count=1

test_sync_pool:
	go test -benchmem -benchtime=10s -bench='^BenchmarkByteBuffer$$' -run=none github.com/smiecj/go_basic/sync_
	go test -benchmem -benchtime=10s -bench='^BenchmarkByteBufferWithPool$$' -run=none github.com/smiecj/go_basic/sync_

test_escape:
	go test -timeout 60s -gcflags="-m" -run ^TestEscapePointer$$ github.com/smiecj/go_basic/basic/escape -v -count=1

test_interview_alphanumber:
	go test -timeout 60s -run ^TestAlphaNumberPrint$$ github.com/smiecj/go_basic/interview/channel -v -count=1

test_interview_tickwithpanic:
	go test -timeout 10s -run ^TestTickWithPanic$$ github.com/smiecj/go_basic/interview/routine -v -count=1

test_interview_waitwithtimeout:
	go test -timeout 60s -run ^TestWaitWithTimeout$$ github.com/smiecj/go_basic/interview/sync_ -v -count=1

test_context_cancel:
	go test -timeout 60s -run ^TestCancelContext$$ github.com/smiecj/go_basic/basic/context -v -count=1

test_routine_leak_send_not_receive:
	go test -timeout 600s -run ^TestChannelSendNotReceive$$ github.com/smiecj/go_basic/interview/routine -v -count=1

test_routine_leak_receive_not_send:
	go test -timeout 600s -run ^TestChannelReceiveNotSend$$ github.com/smiecj/go_basic/interview/routine -v -count=1

test_routine_leak_nil_chan:
	go test -timeout 10s -run ^TestNilChannel$$ github.com/smiecj/go_basic/interview/routine -v -count=1

test_routine_leak_req_no_timeout:
	go test -timeout 30s -run ^TestNetCallWithoutTimeout$$ github.com/smiecj/go_basic/interview/routine -v -count=1

test_routine_leak_lock_without_unlock:
	go test -timeout 30s -race -run ^TestLockWithoutUnlock$$ github.com/smiecj/go_basic/interview/routine -v -count=1

test_routine_leak_wg_not_done:
	go test -timeout 30s -run ^TestWaitGroupNotAllDone$$ github.com/smiecj/go_basic/interview/routine -v -count=1

test_routine_pool_tunny:
	go test -timeout 60s -run ^TestTunnyPool$$ github.com/smiecj/go_basic/basic/routine -v -count=1

test_routine_pool_ants:
	go test -timeout 60s -run ^TestAntsPool$$ github.com/smiecj/go_basic/basic/routine -v -count=1

test_client:
	go test -timeout 60s -run ^TestRequestWithSockProxy$$ github.com/smiecj/go_basic/http -v -count=1

test_fuzz:
	go test -fuzz ^FuzzReverse$$ github.com/smiecj/go_basic/fuzz -v -count=1 -fuzztime 30s

test_mongo:
	go test -timeout 60s -run ^TestMongo$$ github.com/smiecj/go_basic/db/mongo -v -count=1 -mongo=${MONGO_SERVER}

test_zk:
	go test -timeout 60s -run ^TestConnectServer$$ github.com/smiecj/go_basic/backend/zookeeper -v -count=1 -zk=${ZOOKEEPER_SERVER}