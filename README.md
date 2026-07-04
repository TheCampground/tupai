# Tupai (Control Plane for Pangolin)

Represent your entire Pangolin configuration as code. From first time setup, to creating orgs, and adding sites.

### Project Milestones
- [ ] Bootstrap Pangolin
  - [x] Root Account Creation via Config
  - [x] Dynamic config env expansion via reflect api
  - [ ] Organization Creation via Config
  - [ ] ???
- [ ] Config seperation
  - [ ] Seperate container info into different yml
  - [ ] Seperate accounts into folder/file
  - [ ] Seperate organizations into folder/independent file
  - [ ] Seperate API logic (if there is any) into seperate file
  - [ ] ???
- [ ] Users
  - [ ] Create user via config
  - [ ] Delete user via config
  - [ ] Update user via config
  - [ ] ???
- [ ] Sites
  - [ ] Site creation via config
  - [ ] Site deletion via config
  - [ ] Update sites via config
  - [ ] ???
- [ ] Resources
  - [ ] Use existing pangolin blueprints but handle API requests
  - [ ] Possibly more configuration with mantienance pages, and other bits
  - [ ] ???

And so on...


### Config
```yml
version: "0.0.1"

container:
  name: pangolin

rootAccount:
  email: "webmaster@example.com"
  password: "t#st%pa$wOrd1"

api:
  url: https://pangolin.brys.me
  createApiKey: true
```
