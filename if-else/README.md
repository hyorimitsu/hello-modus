If Else
---

This is a conditional branching sample using [Modus](https://modus-continens.com/).

## Usage

1. Analyse a Modusfile and checks the predicate kinds.
    ```shell
    modus check .
    ```

3. Print proof tree of a given query.
    ```shell
    # If you want to print the case of passing "local"
    modus proof . 'app("local")'
    # If you want to print the case of passing "prod"
    modus proof . 'app("prod")'
    # If you want to print both cases
    modus proof . 'app(X)'
    ```

4. Build images.
    ```shell
    # If you want to build the case of passing "local"
    modus build --json . 'app("local")'
    # If you want to build the case of passing "prod"
    modus build --json . 'app("prod")'
    # If you want to build both cases
    modus build --json . 'app(X)'
    ```

5. Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE.
    ```shell
    docker tag sha256:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx app
    ```

6. Run the application.
    ```shell
    docker run -p 8080:8080 app
    ```
