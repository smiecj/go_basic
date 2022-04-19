build:
	go build -o main main.go

run:
	./main

test_float:
	go test -timeout 60s -run ^TestFloatSaveNumBasic$$ github.com/smiecj/go_basic/test -v -count=1

test_slice_sub:
	go test -timeout 60s -run ^TestSubSlice$$ github.com/smiecj/go_basic/basic -v -count=1

test_slice_append:
	go test -timeout 60s -run ^TestSliceAppend$$ github.com/smiecj/go_basic/basic -v -count=1

test_map_get:
	go test -timeout 60s -run ^TestMapGet$$ github.com/smiecj/go_basic/basic -v -count=1

test_escape:
	go test -timeout 60s -gcflags="-m" -run ^TestEscapePointer$$ github.com/smiecj/go_basic/basic/escape -v -count=1

test_interview_alphanumber:
	go test -timeout 60s -run ^TestAlphaNumberPrint$$ github.com/smiecj/go_basic/interview/channel -v -count=1

test_interview_tickwithpanic:
	go test -timeout 10s -run ^TestTickWithPanic$$ github.com/smiecj/go_basic/interview/routine -v -count=1

test_interview_waitwithtimeout:
	go test -timeout 60s -run ^TestWaitWithTimeout$$ github.com/smiecj/go_basic/interview/sync_ -v -count=1