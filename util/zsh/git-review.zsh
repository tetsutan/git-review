
zstyle ':vcs_info:git+post-backend:*' hooks git-review-action
function +vi-git-review-action() {
    local is_review
    is_review=$(command git review status 2>/dev/null)
    if [[ "${is_review}" != "" ]]; then
        hook_com[action]="review"
    fi
}

