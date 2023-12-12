# TIL-CLI

TIL 양식을 커맨드라인으로 자동생성 해줍니다. 여러분의 github 원예사를 흉내내보겠습니다.

**이 프로젝트는 아직 알파 단계입니다. 많은 커맨드라인이 미래에 변경될 것입니다.**

TIL 레포 루트 디렉토리에 `TIL-CLI` 실행파일을 놔주세요.

## init

```sh
./TIL-CLI init
```

- TIL을 시작할 때 내려야할 명령입니다. 마크다운을 만들 때 읽을 기준이되는 파일은 `til-config.json`을 만듭니다.
- `.gitignore`도 자동 생성해줍니다.
- 이미 있으면 명령을 중단합니다.

### til-config.json

```json
{
  "current-project": "진행 중인 프로젝트를 입력해주세요. 지금은 {{current-project-start-day}}일차입니다.\n\n",
  "current-project-start-day": "2023-12-09",
  "show-current-project": true,
  "days-without-accident-day": "2023-12-09",
  "days-without-accident": true,
  "days-without-accident-format": "1일1커밋 무사고: {{days-without-accident-day}}일차\n\n",
  "gratification-format": "## 감사일기\n\n1. ???\n\n",
  "gratification-diary": true,
  "todo": true,
  "todo-format": "## todo\n\n- [ ] ???\n\n---\n\n",
  "retro-format": "## 주간 회고\n\n### Liked\n\n-\n\n### Learned\n\n-\n\n### Lacked\n\n-\n\n### Longed(잘하기 위해 필요한 것)\n\n-\n\n### Action Item\n\n- [ ]",
  "draft": { "today": "", "tomorrow": "", "retro": "" }
}
```

`current-project`는 현재 진행 중인 프로젝트 이름을 입력해주시기 바랍니다. 학습은 프로젝트를 진행 중일 때 합니다. `current-project-start-day`는 프로젝트를 시작하는 날짜를 입력합니다. 만약에 `{{current-project-start-day}}`을 사용하면 시작일과 현재일 사이 차를 구합니다. 프로젝트는 의도적으로 1개만 하도록 만들었습니다. 중요한 것에 집중하게 만들고자 했습니다.

`days-without-accident-format`은 유행하는 무지성 1일1커밋 운동을 실천할 수 있습니다. 커밋을 며칠 연속으로 했는지 시작일을 지정하고 매일 알아서 카운트해줍니다.

무지성 감사일기입니다. 개발자는 이런 감성적인 행동은 필요 없습니다. 멘탈관리에 좋은 활동은 맞습니다.

`todo-format`은 고정적으로 하고 싶은 목록을 작성할 활용할 수 있습니다.

`retro-format`은 회고할 때 고정적으로 활용하고 싶은 포멧을 지원합니다.

## draft 명령

```sh
./TIL-CLI draft
```

실제 TIL 템플릿을 만드는 명령입니다. `til-config.json`을 읽고 오늘 날짜에 해당하는 마크다운을 자동으로 생성합니다.

현재 TIL 마크다운의 파일이름은 다음 형식으로 자동생성합니다.

```
2310/TIL231028.md
```

폴더 구분은 월단위로 했습니다. 프로젝트 단위 구분 기준은 나중에 지원하겠습니다.

```sh
./TIL-CLI draft retro
```

위는 회고 템플릿 생성 명령입니다.

```
2310/TIL231028RetroW.md
```

회고는 기본적으로 주단위를 제공합니다. 월단위 분기단위도 추가 플래그로 지원합니다.

```sh
./TIL-CLI draft retro m # 2310/TIL231028RetroW.md

./TIL-CLI draft retro q # 2310/TIL231028RetroQ.md
```

기본은 오늘로 설정되어 있지만 내일 이번주 일요일 생성 명령도 가능합니다.

```sh
./TIL-CLI draft tomorrow # 2310/TIL231029.md
```

```sh
./TIL-CLI draft sun # 2310/TIL231029.md
```

플래그를 섞으면 아래와 같은 응용도 가능합니다.

```sh
./TIL-CLI draft tomorrow retro # 2310/TIL231029RetroW.md
./TIL-CLI draft sun retro # 2310/TIL231029RetroW.md
```

<!--
매일 작성한 기준 템플릿을 생성합니다.

`template.md`를 생성합니다. 이 파일을 기준으로 TIL을 생성해줄 것입니다.

## today 명령

```sh
./TIL-CLI today
```

temp을 읽고 TIL을 만듭니다.

현재 TIL은 다음 양식으로 자동생성합니다.

```sh
2310/TIL231028.md
```

폴더 구분은 월단위로 했습니다. 프로젝트 단위 구분 기준은 나중에 지원하겠습니다.

~~ 이제 여기서 바로 커밋을 해주세요. ~~

## tomorrow 명령

```sh
./TIL-CLI tomorrow
```

temp을 읽고 TIL을 만듭니다.

현재 TIL은 temp 양식으로 자동생성합니다.

```
2310/
  TIL231027.md
  TIL231028.md
```

폴더 구분은 월단위로 했습니다. 프로젝트 단위 구분 기준은 나중에 지원하겠습니다. ~~참고로 나중이 언제 될지는 모릅니다. ~~

내일 뭘 배울지 아니면 뭘할지 계획도 해보세요. -->
