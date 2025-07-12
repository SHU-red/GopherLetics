# Changelog

## [1.14.2](https://github.com/SHU-red/GopherLetics/compare/v1.14.1...v1.14.2) (2025-07-12)


### Bug Fixes

* **artifact:** Fix GitHub Actions Artifact handling ([c41a71a](https://github.com/SHU-red/GopherLetics/commit/c41a71a9440fad02d3a03f77dd504bf024103af7))

## [1.14.1](https://github.com/SHU-red/GopherLetics/compare/v1.14.0...v1.14.1) (2025-07-12)


### Bug Fixes

* **clean:** Cleanup ([8da63cf](https://github.com/SHU-red/GopherLetics/commit/8da63cf2d597ceba1a011e97d6890cfc0e8c0708))

## [1.14.0](https://github.com/SHU-red/GopherLetics/compare/v1.13.3...v1.14.0) (2025-07-12)


### Features

* Add ls -R to packaging steps for debugging ([837ff6a](https://github.com/SHU-red/GopherLetics/commit/837ff6a66b253f4f53456757c7df7ecfe1232dda))

## [1.13.3](https://github.com/SHU-red/GopherLetics/compare/v1.13.2...v1.13.3) (2025-07-12)


### Bug Fixes

* **clean:** Cleanup ([2cf5d34](https://github.com/SHU-red/GopherLetics/commit/2cf5d34626b6a374e5dcf7f6cb3c17b7c3db4860))

## [1.13.2](https://github.com/SHU-red/GopherLetics/compare/v1.13.1...v1.13.2) (2025-07-11)


### Bug Fixes

* Correct artifact paths for Linux and Darwin builds ([88fa49a](https://github.com/SHU-red/GopherLetics/commit/88fa49a520de24c32d3e9678b957ceb52be3c4b1))

## [1.13.1](https://github.com/SHU-red/GopherLetics/compare/v1.13.0...v1.13.1) (2025-07-11)


### Bug Fixes

* Add app-id to fyne package for Darwin builds ([28389b8](https://github.com/SHU-red/GopherLetics/commit/28389b8d7d4b6ea6f35bc064fbf1325d2426f0a2))

## [1.13.0](https://github.com/SHU-red/GopherLetics/compare/v1.12.12...v1.13.0) (2025-07-11)


### Features

* Add new GitHub Actions workflow for Fyne app build and packaging ([25890ac](https://github.com/SHU-red/GopherLetics/commit/25890acafe91893bb6647b5bd053888de83e9b21))

## [1.12.12](https://github.com/SHU-red/GopherLetics/compare/v1.12.11...v1.12.12) (2025-07-11)


### Bug Fixes

* Remove invalid -pre-build flag from fyne-cross command ([a9fb386](https://github.com/SHU-red/GopherLetics/commit/a9fb386315f60b97ecc8579955148c5fa7dc9a20))

## [1.12.11](https://github.com/SHU-red/GopherLetics/compare/v1.12.10...v1.12.11) (2025-07-11)


### Bug Fixes

* Install libasound2-dev inside fyne-cross container for Linux build ([7b716bc](https://github.com/SHU-red/GopherLetics/commit/7b716bcbb17b495ad9f0d70eb53387a5b06520a4))

## [1.12.10](https://github.com/SHU-red/GopherLetics/compare/v1.12.9...v1.12.10) (2025-07-10)


### Bug Fixes

* Correct Windows binary name and move command in cross-compile.yml ([0449fcd](https://github.com/SHU-red/GopherLetics/commit/0449fcd072ff5e986fb569aa114355b8cee9769e))

## [1.12.9](https://github.com/SHU-red/GopherLetics/compare/v1.12.8...v1.12.9) (2025-07-10)


### Bug Fixes

* Correct fyne-cross appID flag to app-id ([df0893e](https://github.com/SHU-red/GopherLetics/commit/df0893e320986fad9edc7b2ebfb357f6d87110a5))

## [1.12.8](https://github.com/SHU-red/GopherLetics/compare/v1.12.7...v1.12.8) (2025-07-10)


### Bug Fixes

* Rewrite cross-compile.yml with correct YAML structure ([538c472](https://github.com/SHU-red/GopherLetics/commit/538c4720082fe9bf791ad5627cc25357b1427275))

## [1.12.7](https://github.com/SHU-red/GopherLetics/compare/v1.12.6...v1.12.7) (2025-07-10)


### Bug Fixes

* Correct YAML structure of jobs block in cross-compile.yml ([f2ef3cb](https://github.com/SHU-red/GopherLetics/commit/f2ef3cb967ff1bf87f0527dbfa9e2ae2fcb78f2c))

## [1.12.6](https://github.com/SHU-red/GopherLetics/compare/v1.12.5...v1.12.6) (2025-07-10)


### Bug Fixes

* Remove duplicate job definition in cross-compile.yml ([dca7b6f](https://github.com/SHU-red/GopherLetics/commit/dca7b6ffaaf9f20dc764bc6c197ef066c459fee6))

## [1.12.5](https://github.com/SHU-red/GopherLetics/compare/v1.12.4...v1.12.5) (2025-07-10)


### Bug Fixes

* Separate fyne-cross and mingw-w64 setup, add sudo to apt-get ([9507138](https://github.com/SHU-red/GopherLetics/commit/9507138f864d9245b25ba5d13238eefb9b5c4cb6))

## [1.12.4](https://github.com/SHU-red/GopherLetics/compare/v1.12.3...v1.12.4) (2025-07-10)


### Bug Fixes

* Add appID to fyne-cross and ensure fyne-cross is in PATH ([8822083](https://github.com/SHU-red/GopherLetics/commit/8822083fc58ab063352af4db706ee4620af3214e))

## [1.12.3](https://github.com/SHU-red/GopherLetics/compare/v1.12.2...v1.12.3) (2025-07-10)


### Bug Fixes

* Ensure fyne-cross is in PATH and move binaries for go-release-action ([dc11d4a](https://github.com/SHU-red/GopherLetics/commit/dc11d4a0031a698d3bf469a953fd17ef76448be3))

## [1.12.2](https://github.com/SHU-red/GopherLetics/compare/v1.12.1...v1.12.2) (2025-07-10)


### Bug Fixes

* Correct fyne-cross commands and include mingw-w64 for Windows build. Also includes unintended changes to build_linux.yml ([a13f48a](https://github.com/SHU-red/GopherLetics/commit/a13f48aeca4361b7ac40f8caa971dde096a3c7b6))

## [1.12.1](https://github.com/SHU-red/GopherLetics/compare/v1.12.0...v1.12.1) (2025-07-10)


### Bug Fixes

* Correct fyne-cross build commands in cross-compile workflow ([9e3d6a7](https://github.com/SHU-red/GopherLetics/commit/9e3d6a76884caea051f9ee64bdcb2e669b1bfddf))

## [1.12.0](https://github.com/SHU-red/GopherLetics/compare/v1.11.0...v1.12.0) (2025-07-10)


### Features

* "Hello World! Initial commit including actions and standard commit powered by commitizen ([3866b75](https://github.com/SHU-red/GopherLetics/commit/3866b75491128911ee1df6242c7f618a0671abcc))
* **action:** removed alsamixer ([b492a92](https://github.com/SHU-red/GopherLetics/commit/b492a9267d92b94b807fe900e49dc2856eea6caa))
* **actions:** change for build ([eb170c8](https://github.com/SHU-red/GopherLetics/commit/eb170c84337fb7f9cdc3db6c2d27d5333f7ffec7))
* **actions:** force apt-get installation and add libasound ([f8e4f4c](https://github.com/SHU-red/GopherLetics/commit/f8e4f4c03bed555bbc4c8500d12cfbe98172792c))
* **actions:** remove alsa alsa-base etc ([86202d3](https://github.com/SHU-red/GopherLetics/commit/86202d36c6b5ba4dbcb41a7ccf6a2c72629c24dd))
* Add Windows and macOS builds to GitHub Actions ([7bc97f0](https://github.com/SHU-red/GopherLetics/commit/7bc97f03cd042df8a2aeefa7780873a2e54784d1))
* Add Windows and macOS builds to GitHub Actions ([8fd3901](https://github.com/SHU-red/GopherLetics/commit/8fd390158b460542daf5309335ffe08b582269ad))
* **commitizen:** Added commitizen config ([9c66bb0](https://github.com/SHU-red/GopherLetics/commit/9c66bb08687e098c908cf9ef223e34731600ccaf))
* Configure release-please to only run on main branch ([fd467d5](https://github.com/SHU-red/GopherLetics/commit/fd467d50396bf3555c14b44e616d1c09f3bbce37))
* **gui:** Add ability to Play/Pause ([fa9d485](https://github.com/SHU-red/GopherLetics/commit/fa9d48568fb259affe5baff36f6ac20a781bb024))
* Improve README and configuration handling ([66c7fb1](https://github.com/SHU-red/GopherLetics/commit/66c7fb150184a9ae3967a20818b8b87091edbf58))
* **settings:** extend settings struct ([b18cad8](https://github.com/SHU-red/GopherLetics/commit/b18cad845dbb4aa1c11f717b0b287e8f71f1c06c))
* **tts:** Use temp folder for generated audiofiles ([9230086](https://github.com/SHU-red/GopherLetics/commit/923008687347aba6250fa406c08933c8398bfe49))
* Update cross-compile workflow to trigger on release and attach assets ([0b34ca4](https://github.com/SHU-red/GopherLetics/commit/0b34ca402b3fdbe7d16db326f73dacf99ebfa4c4))
* **update:** Update 3rd party packages fyne, toml and zap ([2533a4e](https://github.com/SHU-red/GopherLetics/commit/2533a4e7aeb132bfa28815413041269d08cab460))
* **viper:** Start using viper for configuration reads and writes ([d07fdf9](https://github.com/SHU-red/GopherLetics/commit/d07fdf9f0e3ea3ce984702128c2711c5b85bc9dd))
* **workout fetch:** Added fetching workout & showing as list in GUI ([51cef68](https://github.com/SHU-red/GopherLetics/commit/51cef688ce60363d9c8c60f55b05eb651f77cbc5))


### Bug Fixes

* **actions:** installing alsa ([04c0feb](https://github.com/SHU-red/GopherLetics/commit/04c0feb378b21c6b18393f90dfc7ce5e01692ecb))
* **actions:** re-push for re-build ([dc7c6c3](https://github.com/SHU-red/GopherLetics/commit/dc7c6c38f9efc6e84eaaa99d36462fec31bc91cd))
* **actions:** re-trigger build after actions ([7608e00](https://github.com/SHU-red/GopherLetics/commit/7608e0094191c6b7c743db3e6c7c255a75fb6b19))
* add libegl1-mesa-dev for Wayland build ([f109ae0](https://github.com/SHU-red/GopherLetics/commit/f109ae0e0de8304e2dd4dd0cd43a4ffde8e49aca))
* add Linux-Wayland and Windows build to GitHub action ([d968df2](https://github.com/SHU-red/GopherLetics/commit/d968df2449a41f41a8279761f21169fc81577141))
* adding libxcursor-dev to pre-build ([a68546d](https://github.com/SHU-red/GopherLetics/commit/a68546d134f20a6d2c53e762f93870a9df986560))
* adding libxkbcommon-dev ([2c76689](https://github.com/SHU-red/GopherLetics/commit/2c7668932d0cdf7f5db530da9736708b683fcb9b))
* adding libxrandr-dev ([b97d2b2](https://github.com/SHU-red/GopherLetics/commit/b97d2b283a27dc7f2383b007c2c91e06dd045062))
* adding packages to install for wayland ([0059daa](https://github.com/SHU-red/GopherLetics/commit/0059daaec31f9f2cdf4d90c4bde6d47f068de081))
* also use PAT for release-please ([9e985bd](https://github.com/SHU-red/GopherLetics/commit/9e985bd18adfb3b493f249b2d1d793e53b5cbabb))
* **brower:** Fix re-opening already opened search query ([22275b0](https://github.com/SHU-red/GopherLetics/commit/22275b0c95595ede3a5d4d564f367cc4731028a6))
* **browser:** Prevent opening the same workout in browser ([529944f](https://github.com/SHU-red/GopherLetics/commit/529944f0c04c3ef4385c5e1d18c17c8d38e8cff7))
* **build_linux:** Changed apt-get to apt and added libsdl2-dev ([a1bb8c0](https://github.com/SHU-red/GopherLetics/commit/a1bb8c0aebba99d2bd8cce0b8c4d14620271ae65))
* **build:** trigger ([265d6c7](https://github.com/SHU-red/GopherLetics/commit/265d6c710d3ee1082033f788ebd87949d6d3c9c7))
* **build:** trigger release ([6c0c1f7](https://github.com/SHU-red/GopherLetics/commit/6c0c1f7b08cc9b15148f3966a7ca50a1f84c1f12))
* **buildup:** Implemented tts, formatting of workouts ([3004bab](https://github.com/SHU-red/GopherLetics/commit/3004baba0823fe87d72391b7c0af127d19d64369))
* **buildup:** Introduce Browser search and fix Countdown ([c632109](https://github.com/SHU-red/GopherLetics/commit/c632109c8d0707afde39f01d2f648a0cc7f00895))
* **buildup:** Switch to Go1.22 and Fyne1.3.4 Concurrentcy update ([68a41a9](https://github.com/SHU-red/GopherLetics/commit/68a41a9e302004141431bc8976a1dc45118e9578))
* changed build action trigger to release.published ([29a0616](https://github.com/SHU-red/GopherLetics/commit/29a06163af9710b5d7599a7b86eb1455d2279abc))
* clean Install packages for ubuntu wayland build ([23d22dd](https://github.com/SHU-red/GopherLetics/commit/23d22dd652d1f5907125bbb18621788a5a100537))
* cleaned windows build ([24917cf](https://github.com/SHU-red/GopherLetics/commit/24917cfb6c2ae399740f0b111ad4b8517fd634c3))
* delete bad lines from example for build workflow ([87e01b9](https://github.com/SHU-red/GopherLetics/commit/87e01b988da57e4dcd04540da237bf094114c4b0))
* **execution:** Fixed not correct ending after last workout ([df9a0c3](https://github.com/SHU-red/GopherLetics/commit/df9a0c3300d46f73636652528df4ace87d367312))
* fix release.yaml to build via Action ([b5bfd74](https://github.com/SHU-red/GopherLetics/commit/b5bfd742e0fd4a47ee9da771d874aca2b96d323d))
* fixed Build instructions ([82f51c8](https://github.com/SHU-red/GopherLetics/commit/82f51c85a12abed1023a187608878e09e5f71901))
* **go.mod:** fixed mod file ([9caa997](https://github.com/SHU-red/GopherLetics/commit/9caa997a0fcf06756bb6d3b2e2043d48d96cfb36))
* **gui:** Change appearance of Play/Pause Button on ([fa9d485](https://github.com/SHU-red/GopherLetics/commit/fa9d48568fb259affe5baff36f6ac20a781bb024))
* **gui:** Fix not highlighted Rest sections ([84d27a3](https://github.com/SHU-red/GopherLetics/commit/84d27a3eb4465eeb8a7d826c3cfc8f5cafcdcc86))
* **gui:** UI improvements ([d029a87](https://github.com/SHU-red/GopherLetics/commit/d029a87fdc0484c091f3ac7209b9ae71119848a2))
* ignore missing packages ([abb28bb](https://github.com/SHU-red/GopherLetics/commit/abb28bb65c175cac4d94bd412b2010f12495708b))
* install dependencies as defined at fyne ([537d365](https://github.com/SHU-red/GopherLetics/commit/537d36570751f5a64cc0221c0dbb75ba19f49d45))
* little fix and adding release_tag ([b622e38](https://github.com/SHU-red/GopherLetics/commit/b622e385fb5f3fa965f1cf0c370f71a2cae467ad))
* lksjdf ([7e3710a](https://github.com/SHU-red/GopherLetics/commit/7e3710a288604c0e8b4484cc568aff7c7295da69))
* minor change to test build via action ([b0a052e](https://github.com/SHU-red/GopherLetics/commit/b0a052ed5cd839a75315c314368756e227d9d390))
* no install recommends ([3e331e8](https://github.com/SHU-red/GopherLetics/commit/3e331e802f72ecb8870db2b7aaf22fbdba693f97))
* not specify go version ([905e700](https://github.com/SHU-red/GopherLetics/commit/905e700007475ab4f867a53807d2ebec19c09887))
* only build for linux ([56763ac](https://github.com/SHU-red/GopherLetics/commit/56763ac47cc5d3f2f779b4213cbc87e90e2ed992))
* push for build-Testing reasons ([bb8bd4a](https://github.com/SHU-red/GopherLetics/commit/bb8bd4a8cdc5add7744a38490ffe8e708a2f0496))
* pushing a fix just for testing reasons ([ca263c3](https://github.com/SHU-red/GopherLetics/commit/ca263c3466f8539fbb6836128cf5e4d4fbbb287d))
* Remove hardcoded last-release-ref from release-please ([ae64cc7](https://github.com/SHU-red/GopherLetics/commit/ae64cc71a3c32699aa6b0c043b3820689908eca7))
* remove release-tags ([ee4ad43](https://github.com/SHU-red/GopherLetics/commit/ee4ad4317b895e9c0d963b88d620d434c23fe9c9))
* remove sudo from precommand ([ffdef6e](https://github.com/SHU-red/GopherLetics/commit/ffdef6e3d1337cbb92a7de3677ac2de8e77492ac))
* rename bulid yaml to yml ([9f2d7c8](https://github.com/SHU-red/GopherLetics/commit/9f2d7c8936d45dfa29bf7b10a68550f17199187c))
* **runloop:** Fixing combination of events ([316d75f](https://github.com/SHU-red/GopherLetics/commit/316d75f09c0780fbf28a1ee0ae8c0aefb759bd4a))
* Set release-please last-release-ref to v1.9.0 ([1553d84](https://github.com/SHU-red/GopherLetics/commit/1553d84bb007097aa97aff6303b61ebab96430d7))
* switching to action-gh for build ([25085cd](https://github.com/SHU-red/GopherLetics/commit/25085cde3fb271d8c7558e306c889a6cf45a32f9))
* test ([8568efe](https://github.com/SHU-red/GopherLetics/commit/8568efe3304a72c9f7174d91ebe27e2a3ae41bf3))
* test to get automatic build to work ([d27635f](https://github.com/SHU-red/GopherLetics/commit/d27635f984b2904f7d440f9004de694cccc4ed07))
* test-commit for release-build testing ([ae41e73](https://github.com/SHU-red/GopherLetics/commit/ae41e735c1282612f83cbc26d3fd428e8c3772c3))
* test-Commit to get actions-build working ([9a3d5c9](https://github.com/SHU-red/GopherLetics/commit/9a3d5c9b2484a290f1dad460a6eff23d44ed11d6))
* try only to build linux amd ([7853f99](https://github.com/SHU-red/GopherLetics/commit/7853f997039ad0ac28c912b189bbb2eb60c5e388))
* trying to install libx11-dev during build ([108f2d1](https://github.com/SHU-red/GopherLetics/commit/108f2d1dcd9cd38672f2a9f2d79906b6005e99ee))
* Update release-please action to non-deprecated version ([57b8eca](https://github.com/SHU-red/GopherLetics/commit/57b8eca855bd39649fe88d823549d3d682b4f6c2))
* used advanced build example from go release binaries ([0eee9c9](https://github.com/SHU-red/GopherLetics/commit/0eee9c9978469b254853d383829377da9cc8a905))
* wildcard install libx libraries ([9c7df2a](https://github.com/SHU-red/GopherLetics/commit/9c7df2a35b513aa56183d3afc03f82760319f82e))

## [1.11.0](https://github.com/SHU-red/GopherLetics/compare/v1.10.0...v1.11.0) (2025-07-10)


### Features

* Update cross-compile workflow to trigger on release and attach assets ([0b34ca4](https://github.com/SHU-red/GopherLetics/commit/0b34ca402b3fdbe7d16db326f73dacf99ebfa4c4))

## [1.10.0](https://github.com/SHU-red/GopherLetics/compare/v1.9.0...v1.10.0) (2025-07-10)


### Features

* "Hello World! Initial commit including actions and standard commit powered by commitizen ([3866b75](https://github.com/SHU-red/GopherLetics/commit/3866b75491128911ee1df6242c7f618a0671abcc))
* **action:** removed alsamixer ([b492a92](https://github.com/SHU-red/GopherLetics/commit/b492a9267d92b94b807fe900e49dc2856eea6caa))
* **actions:** change for build ([eb170c8](https://github.com/SHU-red/GopherLetics/commit/eb170c84337fb7f9cdc3db6c2d27d5333f7ffec7))
* **actions:** force apt-get installation and add libasound ([f8e4f4c](https://github.com/SHU-red/GopherLetics/commit/f8e4f4c03bed555bbc4c8500d12cfbe98172792c))
* **actions:** remove alsa alsa-base etc ([86202d3](https://github.com/SHU-red/GopherLetics/commit/86202d36c6b5ba4dbcb41a7ccf6a2c72629c24dd))
* Add Windows and macOS builds to GitHub Actions ([7bc97f0](https://github.com/SHU-red/GopherLetics/commit/7bc97f03cd042df8a2aeefa7780873a2e54784d1))
* Add Windows and macOS builds to GitHub Actions ([8fd3901](https://github.com/SHU-red/GopherLetics/commit/8fd390158b460542daf5309335ffe08b582269ad))
* **commitizen:** Added commitizen config ([9c66bb0](https://github.com/SHU-red/GopherLetics/commit/9c66bb08687e098c908cf9ef223e34731600ccaf))
* Configure release-please to only run on main branch ([fd467d5](https://github.com/SHU-red/GopherLetics/commit/fd467d50396bf3555c14b44e616d1c09f3bbce37))
* **gui:** Add ability to Play/Pause ([fa9d485](https://github.com/SHU-red/GopherLetics/commit/fa9d48568fb259affe5baff36f6ac20a781bb024))
* Improve README and configuration handling ([66c7fb1](https://github.com/SHU-red/GopherLetics/commit/66c7fb150184a9ae3967a20818b8b87091edbf58))
* **settings:** extend settings struct ([b18cad8](https://github.com/SHU-red/GopherLetics/commit/b18cad845dbb4aa1c11f717b0b287e8f71f1c06c))
* **tts:** Use temp folder for generated audiofiles ([9230086](https://github.com/SHU-red/GopherLetics/commit/923008687347aba6250fa406c08933c8398bfe49))
* **update:** Update 3rd party packages fyne, toml and zap ([2533a4e](https://github.com/SHU-red/GopherLetics/commit/2533a4e7aeb132bfa28815413041269d08cab460))
* **viper:** Start using viper for configuration reads and writes ([d07fdf9](https://github.com/SHU-red/GopherLetics/commit/d07fdf9f0e3ea3ce984702128c2711c5b85bc9dd))
* **workout fetch:** Added fetching workout & showing as list in GUI ([51cef68](https://github.com/SHU-red/GopherLetics/commit/51cef688ce60363d9c8c60f55b05eb651f77cbc5))


### Bug Fixes

* **actions:** installing alsa ([04c0feb](https://github.com/SHU-red/GopherLetics/commit/04c0feb378b21c6b18393f90dfc7ce5e01692ecb))
* **actions:** re-push for re-build ([dc7c6c3](https://github.com/SHU-red/GopherLetics/commit/dc7c6c38f9efc6e84eaaa99d36462fec31bc91cd))
* **actions:** re-trigger build after actions ([7608e00](https://github.com/SHU-red/GopherLetics/commit/7608e0094191c6b7c743db3e6c7c255a75fb6b19))
* add libegl1-mesa-dev for Wayland build ([f109ae0](https://github.com/SHU-red/GopherLetics/commit/f109ae0e0de8304e2dd4dd0cd43a4ffde8e49aca))
* add Linux-Wayland and Windows build to GitHub action ([d968df2](https://github.com/SHU-red/GopherLetics/commit/d968df2449a41f41a8279761f21169fc81577141))
* adding libxcursor-dev to pre-build ([a68546d](https://github.com/SHU-red/GopherLetics/commit/a68546d134f20a6d2c53e762f93870a9df986560))
* adding libxkbcommon-dev ([2c76689](https://github.com/SHU-red/GopherLetics/commit/2c7668932d0cdf7f5db530da9736708b683fcb9b))
* adding libxrandr-dev ([b97d2b2](https://github.com/SHU-red/GopherLetics/commit/b97d2b283a27dc7f2383b007c2c91e06dd045062))
* adding packages to install for wayland ([0059daa](https://github.com/SHU-red/GopherLetics/commit/0059daaec31f9f2cdf4d90c4bde6d47f068de081))
* also use PAT for release-please ([9e985bd](https://github.com/SHU-red/GopherLetics/commit/9e985bd18adfb3b493f249b2d1d793e53b5cbabb))
* **brower:** Fix re-opening already opened search query ([22275b0](https://github.com/SHU-red/GopherLetics/commit/22275b0c95595ede3a5d4d564f367cc4731028a6))
* **browser:** Prevent opening the same workout in browser ([529944f](https://github.com/SHU-red/GopherLetics/commit/529944f0c04c3ef4385c5e1d18c17c8d38e8cff7))
* **build_linux:** Changed apt-get to apt and added libsdl2-dev ([a1bb8c0](https://github.com/SHU-red/GopherLetics/commit/a1bb8c0aebba99d2bd8cce0b8c4d14620271ae65))
* **build:** trigger ([265d6c7](https://github.com/SHU-red/GopherLetics/commit/265d6c710d3ee1082033f788ebd87949d6d3c9c7))
* **build:** trigger release ([6c0c1f7](https://github.com/SHU-red/GopherLetics/commit/6c0c1f7b08cc9b15148f3966a7ca50a1f84c1f12))
* **buildup:** Implemented tts, formatting of workouts ([3004bab](https://github.com/SHU-red/GopherLetics/commit/3004baba0823fe87d72391b7c0af127d19d64369))
* **buildup:** Introduce Browser search and fix Countdown ([c632109](https://github.com/SHU-red/GopherLetics/commit/c632109c8d0707afde39f01d2f648a0cc7f00895))
* **buildup:** Switch to Go1.22 and Fyne1.3.4 Concurrentcy update ([68a41a9](https://github.com/SHU-red/GopherLetics/commit/68a41a9e302004141431bc8976a1dc45118e9578))
* changed build action trigger to release.published ([29a0616](https://github.com/SHU-red/GopherLetics/commit/29a06163af9710b5d7599a7b86eb1455d2279abc))
* clean Install packages for ubuntu wayland build ([23d22dd](https://github.com/SHU-red/GopherLetics/commit/23d22dd652d1f5907125bbb18621788a5a100537))
* cleaned windows build ([24917cf](https://github.com/SHU-red/GopherLetics/commit/24917cfb6c2ae399740f0b111ad4b8517fd634c3))
* delete bad lines from example for build workflow ([87e01b9](https://github.com/SHU-red/GopherLetics/commit/87e01b988da57e4dcd04540da237bf094114c4b0))
* **execution:** Fixed not correct ending after last workout ([df9a0c3](https://github.com/SHU-red/GopherLetics/commit/df9a0c3300d46f73636652528df4ace87d367312))
* fix release.yaml to build via Action ([b5bfd74](https://github.com/SHU-red/GopherLetics/commit/b5bfd742e0fd4a47ee9da771d874aca2b96d323d))
* fixed Build instructions ([82f51c8](https://github.com/SHU-red/GopherLetics/commit/82f51c85a12abed1023a187608878e09e5f71901))
* **go.mod:** fixed mod file ([9caa997](https://github.com/SHU-red/GopherLetics/commit/9caa997a0fcf06756bb6d3b2e2043d48d96cfb36))
* **gui:** Change appearance of Play/Pause Button on ([fa9d485](https://github.com/SHU-red/GopherLetics/commit/fa9d48568fb259affe5baff36f6ac20a781bb024))
* **gui:** Fix not highlighted Rest sections ([84d27a3](https://github.com/SHU-red/GopherLetics/commit/84d27a3eb4465eeb8a7d826c3cfc8f5cafcdcc86))
* **gui:** UI improvements ([d029a87](https://github.com/SHU-red/GopherLetics/commit/d029a87fdc0484c091f3ac7209b9ae71119848a2))
* ignore missing packages ([abb28bb](https://github.com/SHU-red/GopherLetics/commit/abb28bb65c175cac4d94bd412b2010f12495708b))
* install dependencies as defined at fyne ([537d365](https://github.com/SHU-red/GopherLetics/commit/537d36570751f5a64cc0221c0dbb75ba19f49d45))
* little fix and adding release_tag ([b622e38](https://github.com/SHU-red/GopherLetics/commit/b622e385fb5f3fa965f1cf0c370f71a2cae467ad))
* lksjdf ([7e3710a](https://github.com/SHU-red/GopherLetics/commit/7e3710a288604c0e8b4484cc568aff7c7295da69))
* minor change to test build via action ([b0a052e](https://github.com/SHU-red/GopherLetics/commit/b0a052ed5cd839a75315c314368756e227d9d390))
* no install recommends ([3e331e8](https://github.com/SHU-red/GopherLetics/commit/3e331e802f72ecb8870db2b7aaf22fbdba693f97))
* not specify go version ([905e700](https://github.com/SHU-red/GopherLetics/commit/905e700007475ab4f867a53807d2ebec19c09887))
* only build for linux ([56763ac](https://github.com/SHU-red/GopherLetics/commit/56763ac47cc5d3f2f779b4213cbc87e90e2ed992))
* push for build-Testing reasons ([bb8bd4a](https://github.com/SHU-red/GopherLetics/commit/bb8bd4a8cdc5add7744a38490ffe8e708a2f0496))
* pushing a fix just for testing reasons ([ca263c3](https://github.com/SHU-red/GopherLetics/commit/ca263c3466f8539fbb6836128cf5e4d4fbbb287d))
* Remove hardcoded last-release-ref from release-please ([ae64cc7](https://github.com/SHU-red/GopherLetics/commit/ae64cc71a3c32699aa6b0c043b3820689908eca7))
* remove release-tags ([ee4ad43](https://github.com/SHU-red/GopherLetics/commit/ee4ad4317b895e9c0d963b88d620d434c23fe9c9))
* remove sudo from precommand ([ffdef6e](https://github.com/SHU-red/GopherLetics/commit/ffdef6e3d1337cbb92a7de3677ac2de8e77492ac))
* rename bulid yaml to yml ([9f2d7c8](https://github.com/SHU-red/GopherLetics/commit/9f2d7c8936d45dfa29bf7b10a68550f17199187c))
* **runloop:** Fixing combination of events ([316d75f](https://github.com/SHU-red/GopherLetics/commit/316d75f09c0780fbf28a1ee0ae8c0aefb759bd4a))
* Set release-please last-release-ref to v1.9.0 ([1553d84](https://github.com/SHU-red/GopherLetics/commit/1553d84bb007097aa97aff6303b61ebab96430d7))
* switching to action-gh for build ([25085cd](https://github.com/SHU-red/GopherLetics/commit/25085cde3fb271d8c7558e306c889a6cf45a32f9))
* test ([8568efe](https://github.com/SHU-red/GopherLetics/commit/8568efe3304a72c9f7174d91ebe27e2a3ae41bf3))
* test to get automatic build to work ([d27635f](https://github.com/SHU-red/GopherLetics/commit/d27635f984b2904f7d440f9004de694cccc4ed07))
* test-commit for release-build testing ([ae41e73](https://github.com/SHU-red/GopherLetics/commit/ae41e735c1282612f83cbc26d3fd428e8c3772c3))
* test-Commit to get actions-build working ([9a3d5c9](https://github.com/SHU-red/GopherLetics/commit/9a3d5c9b2484a290f1dad460a6eff23d44ed11d6))
* try only to build linux amd ([7853f99](https://github.com/SHU-red/GopherLetics/commit/7853f997039ad0ac28c912b189bbb2eb60c5e388))
* trying to install libx11-dev during build ([108f2d1](https://github.com/SHU-red/GopherLetics/commit/108f2d1dcd9cd38672f2a9f2d79906b6005e99ee))
* Update release-please action to non-deprecated version ([57b8eca](https://github.com/SHU-red/GopherLetics/commit/57b8eca855bd39649fe88d823549d3d682b4f6c2))
* used advanced build example from go release binaries ([0eee9c9](https://github.com/SHU-red/GopherLetics/commit/0eee9c9978469b254853d383829377da9cc8a905))
* wildcard install libx libraries ([9c7df2a](https://github.com/SHU-red/GopherLetics/commit/9c7df2a35b513aa56183d3afc03f82760319f82e))

## [1.9.0](https://github.com/SHU-red/GopherLetics/compare/v1.8.0...v1.9.0) (2025-07-10)


### Features

* "Hello World! Initial commit including actions and standard commit powered by commitizen ([3866b75](https://github.com/SHU-red/GopherLetics/commit/3866b75491128911ee1df6242c7f618a0671abcc))
* **action:** removed alsamixer ([b492a92](https://github.com/SHU-red/GopherLetics/commit/b492a9267d92b94b807fe900e49dc2856eea6caa))
* **actions:** change for build ([eb170c8](https://github.com/SHU-red/GopherLetics/commit/eb170c84337fb7f9cdc3db6c2d27d5333f7ffec7))
* **actions:** force apt-get installation and add libasound ([f8e4f4c](https://github.com/SHU-red/GopherLetics/commit/f8e4f4c03bed555bbc4c8500d12cfbe98172792c))
* **actions:** remove alsa alsa-base etc ([86202d3](https://github.com/SHU-red/GopherLetics/commit/86202d36c6b5ba4dbcb41a7ccf6a2c72629c24dd))
* Add Windows and macOS builds to GitHub Actions ([7bc97f0](https://github.com/SHU-red/GopherLetics/commit/7bc97f03cd042df8a2aeefa7780873a2e54784d1))
* Add Windows and macOS builds to GitHub Actions ([8fd3901](https://github.com/SHU-red/GopherLetics/commit/8fd390158b460542daf5309335ffe08b582269ad))
* **commitizen:** Added commitizen config ([9c66bb0](https://github.com/SHU-red/GopherLetics/commit/9c66bb08687e098c908cf9ef223e34731600ccaf))
* Configure release-please to only run on main branch ([fd467d5](https://github.com/SHU-red/GopherLetics/commit/fd467d50396bf3555c14b44e616d1c09f3bbce37))
* **gui:** Add ability to Play/Pause ([fa9d485](https://github.com/SHU-red/GopherLetics/commit/fa9d48568fb259affe5baff36f6ac20a781bb024))
* Improve README and configuration handling ([66c7fb1](https://github.com/SHU-red/GopherLetics/commit/66c7fb150184a9ae3967a20818b8b87091edbf58))
* **settings:** extend settings struct ([b18cad8](https://github.com/SHU-red/GopherLetics/commit/b18cad845dbb4aa1c11f717b0b287e8f71f1c06c))
* **tts:** Use temp folder for generated audiofiles ([9230086](https://github.com/SHU-red/GopherLetics/commit/923008687347aba6250fa406c08933c8398bfe49))
* **update:** Update 3rd party packages fyne, toml and zap ([2533a4e](https://github.com/SHU-red/GopherLetics/commit/2533a4e7aeb132bfa28815413041269d08cab460))
* **viper:** Start using viper for configuration reads and writes ([d07fdf9](https://github.com/SHU-red/GopherLetics/commit/d07fdf9f0e3ea3ce984702128c2711c5b85bc9dd))
* **workout fetch:** Added fetching workout & showing as list in GUI ([51cef68](https://github.com/SHU-red/GopherLetics/commit/51cef688ce60363d9c8c60f55b05eb651f77cbc5))


### Bug Fixes

* **actions:** installing alsa ([04c0feb](https://github.com/SHU-red/GopherLetics/commit/04c0feb378b21c6b18393f90dfc7ce5e01692ecb))
* **actions:** re-push for re-build ([dc7c6c3](https://github.com/SHU-red/GopherLetics/commit/dc7c6c38f9efc6e84eaaa99d36462fec31bc91cd))
* **actions:** re-trigger build after actions ([7608e00](https://github.com/SHU-red/GopherLetics/commit/7608e0094191c6b7c743db3e6c7c255a75fb6b19))
* add libegl1-mesa-dev for Wayland build ([f109ae0](https://github.com/SHU-red/GopherLetics/commit/f109ae0e0de8304e2dd4dd0cd43a4ffde8e49aca))
* add Linux-Wayland and Windows build to GitHub action ([d968df2](https://github.com/SHU-red/GopherLetics/commit/d968df2449a41f41a8279761f21169fc81577141))
* adding libxcursor-dev to pre-build ([a68546d](https://github.com/SHU-red/GopherLetics/commit/a68546d134f20a6d2c53e762f93870a9df986560))
* adding libxkbcommon-dev ([2c76689](https://github.com/SHU-red/GopherLetics/commit/2c7668932d0cdf7f5db530da9736708b683fcb9b))
* adding libxrandr-dev ([b97d2b2](https://github.com/SHU-red/GopherLetics/commit/b97d2b283a27dc7f2383b007c2c91e06dd045062))
* adding packages to install for wayland ([0059daa](https://github.com/SHU-red/GopherLetics/commit/0059daaec31f9f2cdf4d90c4bde6d47f068de081))
* also use PAT for release-please ([9e985bd](https://github.com/SHU-red/GopherLetics/commit/9e985bd18adfb3b493f249b2d1d793e53b5cbabb))
* **brower:** Fix re-opening already opened search query ([22275b0](https://github.com/SHU-red/GopherLetics/commit/22275b0c95595ede3a5d4d564f367cc4731028a6))
* **browser:** Prevent opening the same workout in browser ([529944f](https://github.com/SHU-red/GopherLetics/commit/529944f0c04c3ef4385c5e1d18c17c8d38e8cff7))
* **build_linux:** Changed apt-get to apt and added libsdl2-dev ([a1bb8c0](https://github.com/SHU-red/GopherLetics/commit/a1bb8c0aebba99d2bd8cce0b8c4d14620271ae65))
* **build:** trigger ([265d6c7](https://github.com/SHU-red/GopherLetics/commit/265d6c710d3ee1082033f788ebd87949d6d3c9c7))
* **build:** trigger release ([6c0c1f7](https://github.com/SHU-red/GopherLetics/commit/6c0c1f7b08cc9b15148f3966a7ca50a1f84c1f12))
* **buildup:** Implemented tts, formatting of workouts ([3004bab](https://github.com/SHU-red/GopherLetics/commit/3004baba0823fe87d72391b7c0af127d19d64369))
* **buildup:** Introduce Browser search and fix Countdown ([c632109](https://github.com/SHU-red/GopherLetics/commit/c632109c8d0707afde39f01d2f648a0cc7f00895))
* **buildup:** Switch to Go1.22 and Fyne1.3.4 Concurrentcy update ([68a41a9](https://github.com/SHU-red/GopherLetics/commit/68a41a9e302004141431bc8976a1dc45118e9578))
* changed build action trigger to release.published ([29a0616](https://github.com/SHU-red/GopherLetics/commit/29a06163af9710b5d7599a7b86eb1455d2279abc))
* clean Install packages for ubuntu wayland build ([23d22dd](https://github.com/SHU-red/GopherLetics/commit/23d22dd652d1f5907125bbb18621788a5a100537))
* cleaned windows build ([24917cf](https://github.com/SHU-red/GopherLetics/commit/24917cfb6c2ae399740f0b111ad4b8517fd634c3))
* delete bad lines from example for build workflow ([87e01b9](https://github.com/SHU-red/GopherLetics/commit/87e01b988da57e4dcd04540da237bf094114c4b0))
* **execution:** Fixed not correct ending after last workout ([df9a0c3](https://github.com/SHU-red/GopherLetics/commit/df9a0c3300d46f73636652528df4ace87d367312))
* fix release.yaml to build via Action ([b5bfd74](https://github.com/SHU-red/GopherLetics/commit/b5bfd742e0fd4a47ee9da771d874aca2b96d323d))
* fixed Build instructions ([82f51c8](https://github.com/SHU-red/GopherLetics/commit/82f51c85a12abed1023a187608878e09e5f71901))
* **go.mod:** fixed mod file ([9caa997](https://github.com/SHU-red/GopherLetics/commit/9caa997a0fcf06756bb6d3b2e2043d48d96cfb36))
* **gui:** Change appearance of Play/Pause Button on ([fa9d485](https://github.com/SHU-red/GopherLetics/commit/fa9d48568fb259affe5baff36f6ac20a781bb024))
* **gui:** Fix not highlighted Rest sections ([84d27a3](https://github.com/SHU-red/GopherLetics/commit/84d27a3eb4465eeb8a7d826c3cfc8f5cafcdcc86))
* **gui:** UI improvements ([d029a87](https://github.com/SHU-red/GopherLetics/commit/d029a87fdc0484c091f3ac7209b9ae71119848a2))
* ignore missing packages ([abb28bb](https://github.com/SHU-red/GopherLetics/commit/abb28bb65c175cac4d94bd412b2010f12495708b))
* install dependencies as defined at fyne ([537d365](https://github.com/SHU-red/GopherLetics/commit/537d36570751f5a64cc0221c0dbb75ba19f49d45))
* little fix and adding release_tag ([b622e38](https://github.com/SHU-red/GopherLetics/commit/b622e385fb5f3fa965f1cf0c370f71a2cae467ad))
* lksjdf ([7e3710a](https://github.com/SHU-red/GopherLetics/commit/7e3710a288604c0e8b4484cc568aff7c7295da69))
* minor change to test build via action ([b0a052e](https://github.com/SHU-red/GopherLetics/commit/b0a052ed5cd839a75315c314368756e227d9d390))
* no install recommends ([3e331e8](https://github.com/SHU-red/GopherLetics/commit/3e331e802f72ecb8870db2b7aaf22fbdba693f97))
* not specify go version ([905e700](https://github.com/SHU-red/GopherLetics/commit/905e700007475ab4f867a53807d2ebec19c09887))
* only build for linux ([56763ac](https://github.com/SHU-red/GopherLetics/commit/56763ac47cc5d3f2f779b4213cbc87e90e2ed992))
* push for build-Testing reasons ([bb8bd4a](https://github.com/SHU-red/GopherLetics/commit/bb8bd4a8cdc5add7744a38490ffe8e708a2f0496))
* pushing a fix just for testing reasons ([ca263c3](https://github.com/SHU-red/GopherLetics/commit/ca263c3466f8539fbb6836128cf5e4d4fbbb287d))
* remove release-tags ([ee4ad43](https://github.com/SHU-red/GopherLetics/commit/ee4ad4317b895e9c0d963b88d620d434c23fe9c9))
* remove sudo from precommand ([ffdef6e](https://github.com/SHU-red/GopherLetics/commit/ffdef6e3d1337cbb92a7de3677ac2de8e77492ac))
* rename bulid yaml to yml ([9f2d7c8](https://github.com/SHU-red/GopherLetics/commit/9f2d7c8936d45dfa29bf7b10a68550f17199187c))
* **runloop:** Fixing combination of events ([316d75f](https://github.com/SHU-red/GopherLetics/commit/316d75f09c0780fbf28a1ee0ae8c0aefb759bd4a))
* Set release-please last-release-ref to v1.9.0 ([1553d84](https://github.com/SHU-red/GopherLetics/commit/1553d84bb007097aa97aff6303b61ebab96430d7))
* switching to action-gh for build ([25085cd](https://github.com/SHU-red/GopherLetics/commit/25085cde3fb271d8c7558e306c889a6cf45a32f9))
* test ([8568efe](https://github.com/SHU-red/GopherLetics/commit/8568efe3304a72c9f7174d91ebe27e2a3ae41bf3))
* test to get automatic build to work ([d27635f](https://github.com/SHU-red/GopherLetics/commit/d27635f984b2904f7d440f9004de694cccc4ed07))
* test-commit for release-build testing ([ae41e73](https://github.com/SHU-red/GopherLetics/commit/ae41e735c1282612f83cbc26d3fd428e8c3772c3))
* test-Commit to get actions-build working ([9a3d5c9](https://github.com/SHU-red/GopherLetics/commit/9a3d5c9b2484a290f1dad460a6eff23d44ed11d6))
* try only to build linux amd ([7853f99](https://github.com/SHU-red/GopherLetics/commit/7853f997039ad0ac28c912b189bbb2eb60c5e388))
* trying to install libx11-dev during build ([108f2d1](https://github.com/SHU-red/GopherLetics/commit/108f2d1dcd9cd38672f2a9f2d79906b6005e99ee))
* Update release-please action to non-deprecated version ([57b8eca](https://github.com/SHU-red/GopherLetics/commit/57b8eca855bd39649fe88d823549d3d682b4f6c2))
* used advanced build example from go release binaries ([0eee9c9](https://github.com/SHU-red/GopherLetics/commit/0eee9c9978469b254853d383829377da9cc8a905))
* wildcard install libx libraries ([9c7df2a](https://github.com/SHU-red/GopherLetics/commit/9c7df2a35b513aa56183d3afc03f82760319f82e))

## [1.8.0](https://github.com/SHU-red/GopherLetics/compare/v1.7.3...v1.8.0) (2025-07-10)


### Features

* Add Windows and macOS builds to GitHub Actions ([7bc97f0](https://github.com/SHU-red/GopherLetics/commit/7bc97f03cd042df8a2aeefa7780873a2e54784d1))
* Add Windows and macOS builds to GitHub Actions ([8fd3901](https://github.com/SHU-red/GopherLetics/commit/8fd390158b460542daf5309335ffe08b582269ad))
* Configure release-please to only run on main branch ([fd467d5](https://github.com/SHU-red/GopherLetics/commit/fd467d50396bf3555c14b44e616d1c09f3bbce37))
* Improve README and configuration handling ([66c7fb1](https://github.com/SHU-red/GopherLetics/commit/66c7fb150184a9ae3967a20818b8b87091edbf58))
* **settings:** extend settings struct ([b18cad8](https://github.com/SHU-red/GopherLetics/commit/b18cad845dbb4aa1c11f717b0b287e8f71f1c06c))
* **update:** Update 3rd party packages fyne, toml and zap ([2533a4e](https://github.com/SHU-red/GopherLetics/commit/2533a4e7aeb132bfa28815413041269d08cab460))
* **viper:** Start using viper for configuration reads and writes ([d07fdf9](https://github.com/SHU-red/GopherLetics/commit/d07fdf9f0e3ea3ce984702128c2711c5b85bc9dd))


### Bug Fixes

* Update release-please action to non-deprecated version ([57b8eca](https://github.com/SHU-red/GopherLetics/commit/57b8eca855bd39649fe88d823549d3d682b4f6c2))

## [1.7.3](https://github.com/SHU-red/GopherLetics/compare/v1.7.2...v1.7.3) (2024-07-11)


### Bug Fixes

* **build_linux:** Changed apt-get to apt and added libsdl2-dev ([a1bb8c0](https://github.com/SHU-red/GopherLetics/commit/a1bb8c0aebba99d2bd8cce0b8c4d14620271ae65))

## [1.7.2](https://github.com/SHU-red/GopherLetics/compare/v1.7.1...v1.7.2) (2024-07-09)


### Bug Fixes

* **actions:** installing alsa ([04c0feb](https://github.com/SHU-red/GopherLetics/commit/04c0feb378b21c6b18393f90dfc7ce5e01692ecb))

## [1.7.1](https://github.com/SHU-red/GopherLetics/compare/v1.7.0...v1.7.1) (2024-07-09)


### Bug Fixes

* **go.mod:** fixed mod file ([9caa997](https://github.com/SHU-red/GopherLetics/commit/9caa997a0fcf06756bb6d3b2e2043d48d96cfb36))

## [1.7.0](https://github.com/SHU-red/GopherLetics/compare/v1.6.0...v1.7.0) (2024-07-09)


### Features

* **actions:** remove alsa alsa-base etc ([86202d3](https://github.com/SHU-red/GopherLetics/commit/86202d36c6b5ba4dbcb41a7ccf6a2c72629c24dd))

## [1.6.0](https://github.com/SHU-red/GopherLetics/compare/v1.5.0...v1.6.0) (2024-07-09)


### Features

* **action:** removed alsamixer ([b492a92](https://github.com/SHU-red/GopherLetics/commit/b492a9267d92b94b807fe900e49dc2856eea6caa))

## [1.5.0](https://github.com/SHU-red/GopherLetics/compare/v1.4.0...v1.5.0) (2024-07-09)


### Features

* **actions:** force apt-get installation and add libasound ([f8e4f4c](https://github.com/SHU-red/GopherLetics/commit/f8e4f4c03bed555bbc4c8500d12cfbe98172792c))

## [1.4.0](https://github.com/SHU-red/GopherLetics/compare/v1.3.0...v1.4.0) (2024-07-09)


### Features

* **actions:** change for build ([eb170c8](https://github.com/SHU-red/GopherLetics/commit/eb170c84337fb7f9cdc3db6c2d27d5333f7ffec7))

## [1.3.0](https://github.com/SHU-red/GopherLetics/compare/v1.2.0...v1.3.0) (2024-07-09)


### Features

* **commitizen:** Added commitizen config ([9c66bb0](https://github.com/SHU-red/GopherLetics/commit/9c66bb08687e098c908cf9ef223e34731600ccaf))

## [1.2.0](https://github.com/SHU-red/GopherLetics/compare/v1.1.4...v1.2.0) (2024-06-19)


### Features

* **tts:** Use temp folder for generated audiofiles ([9230086](https://github.com/SHU-red/GopherLetics/commit/923008687347aba6250fa406c08933c8398bfe49))

## [1.1.4](https://github.com/SHU-red/GopherLetics/compare/v1.1.3...v1.1.4) (2024-03-23)


### Bug Fixes

* **build:** trigger release ([6c0c1f7](https://github.com/SHU-red/GopherLetics/commit/6c0c1f7b08cc9b15148f3966a7ca50a1f84c1f12))

## [1.1.3](https://github.com/SHU-red/GopherLetics/compare/v1.1.2...v1.1.3) (2024-03-13)


### Bug Fixes

* **actions:** re-push for re-build ([dc7c6c3](https://github.com/SHU-red/GopherLetics/commit/dc7c6c38f9efc6e84eaaa99d36462fec31bc91cd))

## [1.1.2](https://github.com/SHU-red/GopherLetics/compare/v1.1.1...v1.1.2) (2024-03-13)


### Bug Fixes

* **actions:** re-trigger build after actions ([7608e00](https://github.com/SHU-red/GopherLetics/commit/7608e0094191c6b7c743db3e6c7c255a75fb6b19))

## [1.1.1](https://github.com/SHU-red/GopherLetics/compare/v1.1.0...v1.1.1) (2024-03-13)


### Bug Fixes

* **build:** trigger ([265d6c7](https://github.com/SHU-red/GopherLetics/commit/265d6c710d3ee1082033f788ebd87949d6d3c9c7))

## [1.1.0](https://github.com/SHU-red/GopherLetics/compare/v1.0.33...v1.1.0) (2024-03-13)


### Features

* **workout fetch:** Added fetching workout & showing as list in GUI ([51cef68](https://github.com/SHU-red/GopherLetics/commit/51cef688ce60363d9c8c60f55b05eb651f77cbc5))


### Bug Fixes

* **brower:** Fix re-opening already opened search query ([22275b0](https://github.com/SHU-red/GopherLetics/commit/22275b0c95595ede3a5d4d564f367cc4731028a6))
* **browser:** Prevent opening the same workout in browser ([529944f](https://github.com/SHU-red/GopherLetics/commit/529944f0c04c3ef4385c5e1d18c17c8d38e8cff7))
* **buildup:** Implemented tts, formatting of workouts ([3004bab](https://github.com/SHU-red/GopherLetics/commit/3004baba0823fe87d72391b7c0af127d19d64369))
* **buildup:** Introduce Browser search and fix Countdown ([c632109](https://github.com/SHU-red/GopherLetics/commit/c632109c8d0707afde39f01d2f648a0cc7f00895))
* **buildup:** Switch to Go1.22 and Fyne1.3.4 Concurrentcy update ([68a41a9](https://github.com/SHU-red/GopherLetics/commit/68a41a9e302004141431bc8976a1dc45118e9578))
* **execution:** Fixed not correct ending after last workout ([df9a0c3](https://github.com/SHU-red/GopherLetics/commit/df9a0c3300d46f73636652528df4ace87d367312))
* **gui:** Fix not highlighted Rest sections ([84d27a3](https://github.com/SHU-red/GopherLetics/commit/84d27a3eb4465eeb8a7d826c3cfc8f5cafcdcc86))
* **gui:** UI improvements ([d029a87](https://github.com/SHU-red/GopherLetics/commit/d029a87fdc0484c091f3ac7209b9ae71119848a2))
* **runloop:** Fixing combination of events ([316d75f](https://github.com/SHU-red/GopherLetics/commit/316d75f09c0780fbf28a1ee0ae8c0aefb759bd4a))

## [1.0.33](https://github.com/SHU-red/GopherLetics/compare/v1.0.32...v1.0.33) (2024-01-09)


### Bug Fixes

* add libegl1-mesa-dev for Wayland build ([f109ae0](https://github.com/SHU-red/GopherLetics/commit/f109ae0e0de8304e2dd4dd0cd43a4ffde8e49aca))

## [1.0.32](https://github.com/SHU-red/GopherLetics/compare/v1.0.31...v1.0.32) (2024-01-09)


### Bug Fixes

* clean Install packages for ubuntu wayland build ([23d22dd](https://github.com/SHU-red/GopherLetics/commit/23d22dd652d1f5907125bbb18621788a5a100537))

## [1.0.31](https://github.com/SHU-red/GopherLetics/compare/v1.0.30...v1.0.31) (2024-01-09)


### Bug Fixes

* adding packages to install for wayland ([0059daa](https://github.com/SHU-red/GopherLetics/commit/0059daaec31f9f2cdf4d90c4bde6d47f068de081))

## [1.0.30](https://github.com/SHU-red/GopherLetics/compare/v1.0.29...v1.0.30) (2024-01-09)


### Bug Fixes

* adding libxkbcommon-dev ([2c76689](https://github.com/SHU-red/GopherLetics/commit/2c7668932d0cdf7f5db530da9736708b683fcb9b))

## [1.0.29](https://github.com/SHU-red/GopherLetics/compare/v1.0.28...v1.0.29) (2024-01-08)


### Bug Fixes

* remove release-tags ([ee4ad43](https://github.com/SHU-red/GopherLetics/commit/ee4ad4317b895e9c0d963b88d620d434c23fe9c9))

## [1.0.28](https://github.com/SHU-red/GopherLetics/compare/v1.0.27...v1.0.28) (2024-01-08)


### Bug Fixes

* little fix and adding release_tag ([b622e38](https://github.com/SHU-red/GopherLetics/commit/b622e385fb5f3fa965f1cf0c370f71a2cae467ad))

## [1.0.27](https://github.com/SHU-red/GopherLetics/compare/v1.0.26...v1.0.27) (2024-01-08)


### Bug Fixes

* fixed Build instructions ([82f51c8](https://github.com/SHU-red/GopherLetics/commit/82f51c85a12abed1023a187608878e09e5f71901))

## [1.0.26](https://github.com/SHU-red/GopherLetics/compare/v1.0.25...v1.0.26) (2024-01-08)


### Bug Fixes

* cleaned windows build ([24917cf](https://github.com/SHU-red/GopherLetics/commit/24917cfb6c2ae399740f0b111ad4b8517fd634c3))

## [1.0.25](https://github.com/SHU-red/GopherLetics/compare/v1.0.24...v1.0.25) (2024-01-08)


### Bug Fixes

* add Linux-Wayland and Windows build to GitHub action ([d968df2](https://github.com/SHU-red/GopherLetics/commit/d968df2449a41f41a8279761f21169fc81577141))

## [1.0.24](https://github.com/SHU-red/GopherLetics/compare/v1.0.23...v1.0.24) (2024-01-08)


### Bug Fixes

* install dependencies as defined at fyne ([537d365](https://github.com/SHU-red/GopherLetics/commit/537d36570751f5a64cc0221c0dbb75ba19f49d45))

## [1.0.23](https://github.com/SHU-red/GopherLetics/compare/v1.0.22...v1.0.23) (2024-01-08)


### Bug Fixes

* no install recommends ([3e331e8](https://github.com/SHU-red/GopherLetics/commit/3e331e802f72ecb8870db2b7aaf22fbdba693f97))

## [1.0.22](https://github.com/SHU-red/GopherLetics/compare/v1.0.21...v1.0.22) (2024-01-08)


### Bug Fixes

* ignore missing packages ([abb28bb](https://github.com/SHU-red/GopherLetics/commit/abb28bb65c175cac4d94bd412b2010f12495708b))

## [1.0.21](https://github.com/SHU-red/GopherLetics/compare/v1.0.20...v1.0.21) (2024-01-08)


### Bug Fixes

* lksjdf ([7e3710a](https://github.com/SHU-red/GopherLetics/commit/7e3710a288604c0e8b4484cc568aff7c7295da69))

## [1.0.20](https://github.com/SHU-red/GopherLetics/compare/v1.0.19...v1.0.20) (2024-01-08)


### Bug Fixes

* wildcard install libx libraries ([9c7df2a](https://github.com/SHU-red/GopherLetics/commit/9c7df2a35b513aa56183d3afc03f82760319f82e))

## [1.0.19](https://github.com/SHU-red/GopherLetics/compare/v1.0.18...v1.0.19) (2024-01-08)


### Bug Fixes

* adding libxrandr-dev ([b97d2b2](https://github.com/SHU-red/GopherLetics/commit/b97d2b283a27dc7f2383b007c2c91e06dd045062))

## [1.0.18](https://github.com/SHU-red/GopherLetics/compare/v1.0.17...v1.0.18) (2024-01-08)


### Bug Fixes

* adding libxcursor-dev to pre-build ([a68546d](https://github.com/SHU-red/GopherLetics/commit/a68546d134f20a6d2c53e762f93870a9df986560))

## [1.0.17](https://github.com/SHU-red/GopherLetics/compare/v1.0.16...v1.0.17) (2024-01-08)


### Bug Fixes

* trying to install libx11-dev during build ([108f2d1](https://github.com/SHU-red/GopherLetics/commit/108f2d1dcd9cd38672f2a9f2d79906b6005e99ee))

## [1.0.16](https://github.com/SHU-red/GopherLetics/compare/v1.0.15...v1.0.16) (2024-01-08)


### Bug Fixes

* remove sudo from precommand ([ffdef6e](https://github.com/SHU-red/GopherLetics/commit/ffdef6e3d1337cbb92a7de3677ac2de8e77492ac))

## [1.0.15](https://github.com/SHU-red/GopherLetics/compare/v1.0.14...v1.0.15) (2024-01-08)


### Bug Fixes

* test ([8568efe](https://github.com/SHU-red/GopherLetics/commit/8568efe3304a72c9f7174d91ebe27e2a3ae41bf3))

## [1.0.14](https://github.com/SHU-red/GopherLetics/compare/v1.0.13...v1.0.14) (2024-01-08)


### Bug Fixes

* try only to build linux amd ([7853f99](https://github.com/SHU-red/GopherLetics/commit/7853f997039ad0ac28c912b189bbb2eb60c5e388))

## [1.0.13](https://github.com/SHU-red/GopherLetics/compare/v1.0.12...v1.0.13) (2024-01-08)


### Bug Fixes

* only build for linux ([56763ac](https://github.com/SHU-red/GopherLetics/commit/56763ac47cc5d3f2f779b4213cbc87e90e2ed992))

## [1.0.12](https://github.com/SHU-red/GopherLetics/compare/v1.0.11...v1.0.12) (2024-01-08)


### Bug Fixes

* not specify go version ([905e700](https://github.com/SHU-red/GopherLetics/commit/905e700007475ab4f867a53807d2ebec19c09887))

## [1.0.11](https://github.com/SHU-red/GopherLetics/compare/v1.0.10...v1.0.11) (2024-01-08)


### Bug Fixes

* delete bad lines from example for build workflow ([87e01b9](https://github.com/SHU-red/GopherLetics/commit/87e01b988da57e4dcd04540da237bf094114c4b0))

## [1.0.10](https://github.com/SHU-red/GopherLetics/compare/v1.0.9...v1.0.10) (2024-01-08)


### Bug Fixes

* used advanced build example from go release binaries ([0eee9c9](https://github.com/SHU-red/GopherLetics/commit/0eee9c9978469b254853d383829377da9cc8a905))

## [1.0.9](https://github.com/SHU-red/GopherLetics/compare/v1.0.8...v1.0.9) (2024-01-08)


### Bug Fixes

* switching to action-gh for build ([25085cd](https://github.com/SHU-red/GopherLetics/commit/25085cde3fb271d8c7558e306c889a6cf45a32f9))

## [1.0.8](https://github.com/SHU-red/GopherLetics/compare/v1.0.7...v1.0.8) (2024-01-08)


### Bug Fixes

* also use PAT for release-please ([9e985bd](https://github.com/SHU-red/GopherLetics/commit/9e985bd18adfb3b493f249b2d1d793e53b5cbabb))
* minor change to test build via action ([b0a052e](https://github.com/SHU-red/GopherLetics/commit/b0a052ed5cd839a75315c314368756e227d9d390))

## [1.0.7](https://github.com/SHU-red/GopherLetics/compare/v1.0.6...v1.0.7) (2024-01-08)


### Bug Fixes

* rename bulid yaml to yml ([9f2d7c8](https://github.com/SHU-red/GopherLetics/commit/9f2d7c8936d45dfa29bf7b10a68550f17199187c))

## [1.0.6](https://github.com/SHU-red/GopherLetics/compare/v1.0.5...v1.0.6) (2024-01-08)


### Bug Fixes

* changed build action trigger to release.published ([29a0616](https://github.com/SHU-red/GopherLetics/commit/29a06163af9710b5d7599a7b86eb1455d2279abc))

## [1.0.5](https://github.com/SHU-red/GopherLetics/compare/v1.0.4...v1.0.5) (2024-01-08)


### Bug Fixes

* fix release.yaml to build via Action ([b5bfd74](https://github.com/SHU-red/GopherLetics/commit/b5bfd742e0fd4a47ee9da771d874aca2b96d323d))

## [1.0.4](https://github.com/SHU-red/GopherLetics/compare/v1.0.3...v1.0.4) (2024-01-08)


### Bug Fixes

* test-Commit to get actions-build working ([9a3d5c9](https://github.com/SHU-red/GopherLetics/commit/9a3d5c9b2484a290f1dad460a6eff23d44ed11d6))

## [1.0.3](https://github.com/SHU-red/GopherLetics/compare/v1.0.2...v1.0.3) (2024-01-08)


### Bug Fixes

* test-commit for release-build testing ([ae41e73](https://github.com/SHU-red/GopherLetics/commit/ae41e735c1282612f83cbc26d3fd428e8c3772c3))

## [1.0.2](https://github.com/SHU-red/GopherLetics/compare/v1.0.1...v1.0.2) (2024-01-08)


### Bug Fixes

* test to get automatic build to work ([d27635f](https://github.com/SHU-red/GopherLetics/commit/d27635f984b2904f7d440f9004de694cccc4ed07))

## [1.0.1](https://github.com/SHU-red/GopherLetics/compare/v1.0.0...v1.0.1) (2024-01-08)


### Bug Fixes

* push for build-Testing reasons ([bb8bd4a](https://github.com/SHU-red/GopherLetics/commit/bb8bd4a8cdc5add7744a38490ffe8e708a2f0496))
* pushing a fix just for testing reasons ([ca263c3](https://github.com/SHU-red/GopherLetics/commit/ca263c3466f8539fbb6836128cf5e4d4fbbb287d))

## 1.0.0 (2024-01-08)


### Features

* "Hello World! Initial commit including actions and standard commit powered by commitizen ([3866b75](https://github.com/SHU-red/GopherLetics/commit/3866b75491128911ee1df6242c7f618a0671abcc))
