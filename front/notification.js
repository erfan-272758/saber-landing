
var stack_top_left = { "dir1": "down", "dir2": "right", "push": "top" };
function showNotification(header, text, type) {
    new PNotify({
        title: header,
        text:text,
        addclass: 'alert alert-styled-left alert-arrow-left stack-top-right',
        type: type,
        stack: stack_top_left,
        delay: 10000,
    });
}