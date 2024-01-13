A basic rule for writing cryptographic implementations is to wipe any secrets as soon as they're no longer needed.  The longer a secret is kept, the higher the chance that someone
will be able to access it.  For transient secrets, this involves wiping the memory locations.  This has to be done intentionally by the developer as some operating systems don't bother
wiping out the memory before they give it to the next application.
