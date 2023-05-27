$(".body_articles_tag").on('click', function () {
    var tagname = $(this).find("a").text()
    console.log(tagname)
    $.ajax({
        url: 'http://localhost:8080/Gentags',
        method: 'POST',
        data: { tagname: tagname },
        success: function (response) {
            console.log(response)
        }
    })
})