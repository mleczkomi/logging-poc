<?php

declare(strict_types=1);

namespace App\Server;

use Swoole\Http\Server;

/**
 * @author Michał Mleczko <michal.mleczko@iiit.pl>
 */
class Factory
{
    public static function create(): Server
    {
        $server = new Server('0.0.0.0', $_ENV['APP_ENV'] === 'dev' ? 8080 : 80, SWOOLE_BASE);
        $server->set([
            'worker_num' => 4,
            'task_worker_num' => 4,
            'http_parse_post' => false,
            'task_enable_coroutine' => true,
        ]);

        return $server;
    }
}
