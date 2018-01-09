$(function () {
    $('[data-toggle="tooltip"]').tooltip();

    $.each($('#html-gen h1, #html-gen h2, #html-gen h3, #html-gen h4, #html-gen h5, #html-gen h6'), function(index, value) {
        var anchor = $(this).attr("id");

        $(this).html("<a class='anchor' href='#" + anchor + "'>#</a> " + $(this).html());
    });

    setTimeout(function() {
        if (window.location.hash != "") {
            el = document.getElementById(window.location.hash.replace("#", ""));
            if (el !== null) { el.scrollIntoView(); }
        }
    }, 500);
});
