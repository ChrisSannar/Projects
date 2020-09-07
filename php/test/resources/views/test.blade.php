@extends('layouts.app')

@section('content')
<div class="container">
    <div>{{ date("H:m:s") }}</div>
    @foreach([1,2,3,4,5] as $num)
        @if($num % 2 == 1)
            <p>Imma test: {{ $num }}</p>
        @else
            <p>Another test: {{ $num }}</p>
        @endif
    @endforeach
</div>
@endsection