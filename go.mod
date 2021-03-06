module github.com/manifoldco/grafton

require (
	github.com/PuerkitoBio/purell v1.1.0 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/alecthomas/gometalinter v2.0.11+incompatible
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf
	github.com/client9/misspell v0.3.4
	github.com/go-openapi/analysis v0.0.0-20170813233457-8ed83f2ea9f0 // indirect
	github.com/go-openapi/errors v0.0.0-20170426151106-03cfca65330d
	github.com/go-openapi/inflect v0.0.0-20130829110746-b1f6470ffb9c // indirect
	github.com/go-openapi/jsonpointer v0.0.0-20170102174223-779f45308c19 // indirect
	github.com/go-openapi/jsonreference v0.0.0-20161105162150-36d33bfe519e // indirect
	github.com/go-openapi/loads v0.0.0-20170520182102-a80dea3052f0 // indirect
	github.com/go-openapi/runtime v0.0.0-20170303002511-e66a4c440602
	github.com/go-openapi/spec v0.0.0-20170928160009-48c2a7185575 // indirect
	github.com/go-openapi/strfmt v0.0.0-20170822153411-610b6cacdcde
	github.com/go-openapi/swag v0.0.0-20170606142751-f3f9494671f9
	github.com/go-openapi/validate v0.0.0-20170921144055-dc8a684882cf
	github.com/go-swagger/go-swagger v0.0.0-20170414161553-fbc64c262a83
	github.com/go-swagger/scan-repo-boundary v0.0.0-20180623220736-973b3573c013 // indirect
	github.com/go-zoo/bone v0.0.0-20160911183509-fd0aebc74e90
	github.com/gobuffalo/packr v1.13.7
	github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/google/uuid v1.1.0 // indirect
	github.com/gordonklaus/ineffassign v0.0.0-20180909121442-1003c8bd00dc
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/handlers v1.4.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/jessevdk/go-flags v1.4.0 // indirect
	github.com/manifoldco/go-base32 v1.0.2
	github.com/manifoldco/go-base64 v1.0.1
	github.com/manifoldco/go-connector v0.0.2
	github.com/manifoldco/go-jwt v0.1.1
	github.com/manifoldco/go-manifold v0.9.5
	github.com/manifoldco/go-signature v1.0.1
	github.com/manifoldco/promptui v0.3.2
	github.com/onsi/ginkgo v1.7.0 // indirect
	github.com/onsi/gomega v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/sirupsen/logrus v1.0.3
	github.com/spf13/viper v1.2.1 // indirect
	github.com/toqueteos/webbrowser v1.1.0 // indirect
	github.com/tsenart/deadcode v0.0.0-20160724212837-210d2dc333e9
	github.com/tylerb/graceful v1.2.15 // indirect
	github.com/urfave/cli v1.20.0
	golang.org/x/crypto v0.0.0-20181126144156-c05539cddb59
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

// This version of kingpin is incompatible with the released version of
// gometalinter until the next release of gometalinter, and possibly until it
// has go module support, we'll need this exclude, and perhaps some more.
//
// After that point, we should be able to remove it.
exclude gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20180810215634-df19058c872c
