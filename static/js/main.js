$(function () {
    $('[data-toggle="tooltip"]').tooltip();

    $.each($('#html-gen h2, #html-gen h3, #html-gen h4'), function(index, value) {
        var anchor = makeAnchor($(this).text());
        $(this).attr("id", anchor);

        $(this).html("<a class='anchor' href='#" + anchor + "'>#</a> " + $(this).html());
    });

    setTimeout(function() {
        if (window.location.hash != "") {
            el = document.getElementById(window.location.hash.replace("#", ""));
            if (el !== null) { el.scrollIntoView(); }
        }
    }, 500);
});

var anchorArray = [];

function makeAnchor(text) {
    text = text.toLowerCase().replace(/[^a-z0-9-_ ]/g, "").trim().replace(/[-_ ]+/g, "-");

    var count = 0;
    for (i = 0; i < anchorArray.length; i++) {
        if (text == anchorArray[i]) {
            count++
        }
    }

    anchorArray.push(text);

    if (count == 0) { return text; }

    return text + "-" + (count + 1);
}
