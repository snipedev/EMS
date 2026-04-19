async function AuthRoutes(fastify,options) {
    fastify.post('/login', {
        config:{
            rateLimit: {
                max: 5,
                timeWindow: '1 minute'
            }
        }
    }, async (request, reply) => {
        const { username, password } = request.body;
        // Placeholder for authentication logic
        if (username === 'admin' && password === 'password') {
            const token = fastify.jwt.sign({ username });
            //localStorage.setItem('token', token);
            return reply.status(200).send({ "token-generated": token });
        } else {
            reply.status(401).send({ error: 'Invalid credentials' });
        }
    });

}

module.exports = AuthRoutes;