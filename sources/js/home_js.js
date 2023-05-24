$('#search').on('input', function() {
    $.ajax({
        url: '/Search',
        method: 'POST',
        data: {keywords: $(this).val()},
        success: function(results) {
            $('#search-results').html(results) 
        }
    })
}) 