@extends('layouts.app')

@section('content')
<div class="container">
    <form action="/tasks" enctype="multipart/form-data" method="post">
        <!-- Cross Site Request Forgery Protections - Make sure the data is coming from the site -->
        @csrf

        <div class="card-header">
            New Task
        </div>
        <div class="card-body offset-2">
            <div class="form-group row">
                <input name="text" type="text" id="task">
                <button type="submit">Add</button>
            </div>
            @if ($errors->has('text'))
                <p>{{ $errors->first('text') }}</p>
            @endif
            <a href="/tasks">Back</a>
        </div>
    </form>
</div>
@endsection
