<script>
    import { page } from '$app/state';
</script>

<header>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container-fluid">
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNavDropdown">
                <ul class="navbar-nav">
                    <li class="nav-item" aria-current={page.url.pathname === '/' ? 'page' : undefined}>
                        <a class="nav-link active" aria-current="page" href="/">Home</a>
                    </li>
                    <li class="nav-item" aria-current={page.url.pathname === '/about' ? 'page' : undefined}>
                        <a class="nav-link" href="/about">About</a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</header>