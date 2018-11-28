#ifndef BIGINT_H
#define BIGINT_H

#include "define.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct bigint_t bigint_t;
typedef void (*bigint_bitwise_callback_t)(bigint_t*, bigint_t*, bigint_t*, unsigned int);


bigint_t *bigint(int n);
void bigint_free(bigint_t *a);
bigint_t *bigint_clone(bigint_t *a);
bigint_t *bigint_shift_left(bigint_t *a, unsigned int n_shifts);
bigint_t *bigint_shift_right(bigint_t *a, unsigned int n_shifts);
void bigint_set_bit(bigint_t *dst, unsigned int bit_index, unsigned int bit_value);
unsigned int bigint_get_bit(bigint_t *a, unsigned int bit_index);
unsigned int bigint_degree(bigint_t *a);
int bigint_compare(bigint_t *a, bigint_t *b);
bigint_t *bigint_mul(bigint_t *a, bigint_t *b);
bigint_t *bigint_sub(bigint_t *a, bigint_t *b);
bigint_t *bigint_divrem(bigint_t *dividend, bigint_t *divisor, bigint_t **remainder);
bigint_t *bigint_add(bigint_t *a, bigint_t *b);
COIN_API_EXPORT void bigint_to_string(bigint_t *a, unsigned int base, char *dst);
COIN_API_EXPORT bigint_t *bigint_from_string(char *a, unsigned int base);
bigint_t *bigint_factorial(bigint_t *a);
bigint_t *bigint_xor(bigint_t *a, bigint_t *b);
bigint_t *bigint_and(bigint_t *a, bigint_t *b);
bigint_t *bigint_or(bigint_t *a, bigint_t *b);
bigint_t *bigint_not(bigint_t *a);

static void bigint_allocate_block(bigint_t *a, unsigned int n);
static void bigint_xor_bit_mutable(bigint_t *dst, unsigned int bit_index, unsigned int bit_value);
static bigint_t *bigint_apply_bitwise_function(bigint_t *a, bigint_t *b, bigint_bitwise_callback_t callback);
static void bigint_xor_callback(bigint_t *dst, bigint_t *a, bigint_t *b, unsigned int bit_index);
static void bigint_and_callback(bigint_t *dst, bigint_t *a, bigint_t *b, unsigned int bit_index);
static void bigint_or_callback(bigint_t *dst, bigint_t *a, bigint_t *b, unsigned int bit_index);
static bigint_t *bigint_shift_left_fast(bigint_t *a, unsigned int n_shifts);
static bigint_t *bigint_shift_right_fast(bigint_t *a, unsigned int n_shifts);

#ifdef __cplusplus
}
#endif

#endif