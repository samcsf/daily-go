$('#submit-alr-fail').hide();
$('#submit-alr-succ').hide();

var list = [];
var editing = -1;
var queryCategory = function() {
  var $nav = $('#category-list');
  $('#content-text').empty().append('<p>Go Go Go!</p>');
  $.ajax({
    url: '/post',
    dataType: 'json',
  }).done(function(data){
    list = data;
    $nav.empty();
    data.forEach(function(d, i){
      $nav.append('<li class="nav-item"><a class="nav-link" href="#" onclick="showContent(' + i + ')">'+ d.title + '</a></li>')
    });
  });
}
queryCategory();

var showContent = function(i) {
  var $text = $('#content-text');
  $text.empty();
  $text.append('<h3>' + list[i].title + '</h3>');
  $text.append(marked(list[i].content));
  $text.append('<button type="button" class="btn btn-success btn-sm" onclick="onEdit(' + i + ')">Edit</button>');
  $text.append('<button type="button" class="btn btn-danger btn-sm" onclick="onDelete(' + i + ')">Delete</button>');
};

var simplemde = new SimpleMDE({ element: document.getElementById('editor') });

var showSucc = function() {
  $('#submit-alr-succ').show();
  $('#submit-alr-succ').fadeOut(1200);
};
var showFail = function() {
  $('#submit-alr-fail').show();
  $('#submit-alr-fail').fadeOut(1200);
};

var onSubmit = function() {
  var content = simplemde.value();
  var title = $('#input-title').val();
  var method = editing > -1 ? 'PUT': 'POST';
  var json = { title, content };
  if (editing > -1) {
    json.id = list[editing].id;
  }
  $.ajax({
    method,
    url: '/post',
    data: JSON.stringify(json),
  }).done(function(data){
    showSucc();
    simplemde.value('');
    $('#input-title').val('');
    queryCategory();
  }).fail(function(){
    showFail();
  });
}

var onDelete = function(i) {
  $.ajax({
    method: 'DELETE',
    url: '/post',
    data: JSON.stringify(list[i]),
  }).done(function(data){
    showSucc();
    queryCategory();
  }).fail(function(){
    showFail();
  });
}

var onEdit = function(i) {
  editing = i;
  simplemde.value(list[i].content);
  $('#input-title').val(list[i].title); 
}