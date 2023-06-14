# Chatanium Project
## Cross-Flatform Chatbot Runtime
본 프로젝트의 목적은 플랫폼에 구애 받지 않은 편리한 챗봇 개발을 위해 설계된 런타임입니다.
하지만, 채팅 플랫폼은 각자의 메시지 타입과 지원하는 메시지의 종류도 전부 다릅니다. 따라서 우리는 메세지를 정형화 시키기 보다는, 각자의 메신저에 맞춰 모듈을 설계할 수 있도록 선택권을 부여하기로 했습니다.
## Backend Module
Golang 뿐만 아닌 다양한 언어의 라이브러리 생태계는 가늠 할 수 없을 정도로 방대합니다.
따라서 우리는 Backend Module이라는 개념을 도입했습니다.
인터페이스에 맞춰 모듈을 설계하면 자동으로 런타임을 통해 챗봇 모듈에 분배합니다.
**Backend Module을 만들고 싶다면 이곳을 클릭하세요.**
## Chatbot Module
이 런타임의 꽃이라고 할 수 있는 Chatbot Module도 마찬가지로 인터페이스에 맞춰
모듈을 설계하기만 하면 됩니다. 복잡한 크로스-플랫폼 채팅 연동 기능을 더이상 생각하지 마세요.
**Chatbot Module을 만들고 싶다면 이곳을 클릭하세요.**
## Fit for Module Development
당연한 말이지만, 이 런타임은 모듈 개발의 편의를 중점으로 두고 설계 했습니다.
따라서 모듈을 만드는 방법을 다양하게 만들었고, 그 중 하나를 꼽자면 폭 넓은 모듈 지원 범위를 꼽을 수 있겠습니다.
## Various way to Import module
우리는 다양한 모듈 임포트 방식을 지원하기 위해 노력하고 있습니다. 정적 컴파일 방식부터, RPC over HTTPS로 원격 모듈을 폭 넓게 지원 할 수 있도록 전부 개발 계획에 존재하며, 실제로 이 모든 방식은 개발 중에 있습니다.
