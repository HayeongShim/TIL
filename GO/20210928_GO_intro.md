쉽고 빠른 Go 시작하기
==

노마드 코더 무료 강의로 배워보는 Go 언어. 이것저것 재지 않고 일단 시작해보기로 했다.  
</br>
  
## 1. What we are building?
1. kr.indeed.com의 데이터를 추출하는 web scrapper.  
추출한 데이터를 csv 형식으로 볼 수 있도록 한다.  
Go에 있는 데이터 처리 도구(multi-core processing, concurrent)를 사용하여  
우리가 할 수 있는 최대의 속도를 내보자.  
2. 또한, Echo를 활용한 엄청 고성능의 미니멀한 scrapper 웹사이트도 만들어 볼 것.
</br></br>

## 2. Software and Installation  
1. Go를 다운로드한다 (https://golang.org/)  
2. Go PATH는 환경변수인데, 자동적으로 생성된 Go 폴더가 꼭 있어야 한다. 터미널에 **go env**를 입력하여 'GOPATH'의 결과를 보면 된다.
Go는 내가 원하는 디렉토리에 프로젝트를 만들어서 사용할 수 없다.  
Desktop.. Document.. Github.. nope! 무조건 Go PATH 디렉토리에 저장되어야 함.
- **ISSUE** ->  
  나의 경우 GOPATH="/Users/hayeongshim/go" 디렉토리가 없어서 직접 생성 및 하위에 bin, pkg, src 폴더를 만들어주었다.
3. Node.js나 Python은 모듈이나 패키지를 다운받을 수 있는 곳이 한정적ㅡex. npm, pypiㅡ인데, Go에서는 내가 원하는 곳 어디에서든 코드를 다운받아 사용할 수 있다.  
받아온 코드들을 보기좋게 정리하는 방법은, 도메인별로 분류하는 것. github / golang / google etc.  

![screenshot](https://postfiles.pstatic.net/MjAyMTA5MjlfMTQ0/MDAxNjMyODQyNTM5Nzk1.Ac6bLuMlS-Z83LwpBBp58p9akbI81s-643mu8oEu_GQg.KiHNNECo18wqZLXF1dauAfP8MWMTnqv7pzK66kfEPS8g.PNG.shyme2055/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7_2021-09-29_%EC%98%A4%EC%A0%84_12.20.43.png?type=w966)
