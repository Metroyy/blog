$(".tag-item").on('click', function () {
    var tagname = $(this).find("a").text()
    tagname = tagname.substring(0,tagname.length-3)
    tagname = tagname.trim()
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