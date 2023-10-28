# TIL-CLI

TIL 양식을 커맨드라인으로 자동생성 해줍니다. 여러분의 github 원예사를 흉내내보겠습니다.

**이 프로젝트는 아직 알파 단계입니다. 많은 커맨드라인이 미래에 변경될 것입니다.**

TIL 레포 루트 디렉토리에 `TIL-CLI` 실행파일을 놔주세요.

## temp 명령

```sh
./TIL-CLI temp
```

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

내일 뭘 배울지 아니면 뭘할지 계획도 해보세요.