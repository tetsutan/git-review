
# git-review

Reviewing tool for git

![git-review](https://raw.githubusercontent.com/tetsutan/git-review/master/screenshot.gif)


## Install

```
make install
```

## Usage

After `cd your-repo-path`

#### start

```
git review start develop..feature/xxx
```

#### stop reviewing

```
git review reset
```

#### help


```
git review help
```

#### show diff

```
git review difftool  # alias `v`
git review diff  # alias `d`
```


#### status
```
git review status  # alias `s`
```

#### comment

Comment to current review.

```
git review comment lgtm
git review good # alias `comment good` and `next`
git review bad # alias `comment bad` and `next`
git review skip # alias `comment skip` and `next`
```

#### next/prev

Move next(prev) commit without comment

```
git review next
git review prev
```

#### next-u/prev-u

Move to next uncommented commit

```
git review next-u # alias `nu`
git review prev-u # alias `pu`
```

#### complete

Output review log.

```
git review complete
```

## zsh integration

Copy `util/zsh/git-review.zsh` to your zsh configuration

