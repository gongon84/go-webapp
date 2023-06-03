// 次へボタン
function nextPage(event) {
    event.preventDefault();
    const params = new URLSearchParams(window.location.search)
    let startParam = Number(params.get('start'))
    params.set('start', startParam + 1)
    window.location.search = params.toString()
}

// 前へボタン
function beforePage(event) {
    event.preventDefault();
    const params = new URLSearchParams(window.location.search)
    let startParam = Number(params.get('start'))
    params.set('start', startParam - 1)
    window.location.search = params.toString()
}