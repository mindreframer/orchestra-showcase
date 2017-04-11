## Showcase how Orchestra (https://github.com/monzo/orchestra) works


## Installation
    # install orchestra
    $ go get install -u github.com/monzo/orchestra
    $ go get github.com/mindreframer/orchestra-showcase
    $ cd $GOPATH/src/github.com/mindreframer/orchestra-showcase


## Play with orchestra
    # start all services
    $ orchestra start

    # show logs for all
    $ orchestra logs

    # show logs just for srv.07
    $ orchestra logs srv.07

    # stop just srv.07
    $ orchestra stop srv.07

    # show PIDS / status for srvs
    $ orchestra ps

    # stop all services
    $ orchestra stop

