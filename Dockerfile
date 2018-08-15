FROM library/golang

# dep for vendoring
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Recompile the standard library without CGO
# RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/back/beego_api
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./beego_api)
ADD . $APP_DIR

# Compile the binary and statically link
# -ldflags '-w -s'
#   -s: 去掉符号表
#   -w: 去掉调试信息，不能gdb调试了
RUN cd $APP_DIR && go build -ldflags '-w -s'

EXPOSE 8080
