# tabssh

[idk](https://twitter.com/rsnous/status/1365106287080472579)

Uses [TabFS](https://github.com/osnr/TabFS) and
[gilderlabs/ssh](https://github.com/gliderlabs/ssh).

Set your TabFS mount path in `tabssh.go`.

```
$ go run tabssh.go
```

and

```
$ ssh -o StrictHostKeyChecking=no localhost -p 2222
```

(you can set a fun hostname in `~/.config/ssh`:
```
Host last-focused-tab.safari.localhost
    HostName localhost
    Port 2222
    LogLevel ERROR
    StrictHostKeyChecking no
    UserKnownHostsFile /dev/null 
```
)

## ideas

a hack would be to make it dispatch to tab depending on provided username, like how
[their Docker
example](https://github.com/gliderlabs/ssh/tree/master/_examples/ssh-docker)
dispatches to process (ssh `jq@localhost` runs `jq`)

but really, the _right_ way to do it would be to make it so that
`safari.localhost` is a hostname that actually lets you talk to
Safari, and `last-focused-tab.safari.localhost` is a hostname that
actually lets you talk to the last focused tab.

could you make a virtual network or something to do that? (where each
tab is a host on the virtual network) I mean, that feels
philosophically right; [tabs are virtual
computers](https://twitter.com/rsnous/status/1352014584731734016), so
maybe they should be network-addressable like your real computer is
(and maybe other things should be network-addressable that way too --
individual applications, documents, etc).

like 'port numbers' feel kinda like they unnaturally promote one level
of computer, the physical one on your desk that has a Wi-Fi chip, and
hide the computers nested inside it (such as each of your browser
tabs)
