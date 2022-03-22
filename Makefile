test_float:
	go test -timeout 60s -run ^TestFloatSaveNumBasic$$ github.com/smiecj/go_basic/test -v -count=1

test_interview_alphanumber:
		go test -timeout 60s -run ^TestAlphaNumberPrint$$ github.com/smiecj/go_basic/interview/channel -v -count=1